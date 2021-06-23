package neo4j_tentacle

import (
	"LinshanYu/neo4j-go-cypher/connector"
	"LinshanYu/neo4j-go-cypher/params_container"
	"LinshanYu/neo4j-go-cypher/utils"
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"strings"
	"sync"
)

type cypher struct {
	accessMode     neo4j.AccessMode
	header         PrimaryCypher
	this           PrimaryCypher
	lock           *sync.Mutex
	decorateCypher map[decorateName][]DecorateCypher
	conditions     map[string][]Condition
	keys           []string
}

type decorateName string

var returnDecorate decorateName = "return"
var limitDecorate decorateName = "limit"
var orderDecorate decorateName = "order"

func (m *cypher) ReNew(name string, cypherType CypherType, action CypherAction) Cypher {
	m.lock.Lock()
	defer m.lock.Unlock()
	var cypher PrimaryCypher
	switch cypherType {
	case RelationType:
		cypher = &relationCypher{
			ca: action,
		}
	case NodeType:
		cypher = &nodeCypher{
			ca: action,
		}
	}
	if neo4j.AccessModeWrite == cypherActionToAccessMode(action) {
		m.accessMode = neo4j.AccessModeWrite
	}
	m.this.setNext(cypher)
	m.this = m.this.next()
	name = strings.Trim(name, " ")
	if "" == name {
		name = utils.GetRandomString(1) + utils.GetRandomString1(5)
	}
	m.this.name(name)
	return m
}

func (m *cypher) Merge(condition map[string]interface{}) Cypher {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.this.setMerge(condition)
	return m
}

func (m *cypher) Node(node *params_container.Node) Cypher {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.this.setLabel(node.Label)
	m.this.setProperty(node.Property)
	return m

}

func (m *cypher) Relation(fromNodeAlias, toNodeAlias string, relations []string, relationProperty map[string]interface{}, direction bool) Cypher {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.this.setRelationFrom(fromNodeAlias)
	m.this.setRelationTo(toNodeAlias)
	m.this.setRelations(relations, direction)
	m.this.setProperty(relationProperty)
	return m
}

func (m *cypher) Return(distinct bool) Cypher {
	m.lock.Lock()
	defer m.lock.Unlock()
	r := &returnCypher{
		distinct: distinct,
		nodeName: m.this.getName(),
	}
	m.keys = append(m.keys, m.this.getName())
	m.decorateCypher[returnDecorate] = append(m.decorateCypher[returnDecorate], r)
	return m
}

func (m *cypher) OrderBy(propertyName string, asc bool) Cypher {
	m.lock.Lock()
	defer m.lock.Unlock()
	d := &orderCypher{
		nodeName: m.this.getName(),
		order:    propertyName,
		asc:      asc,
	}
	m.decorateCypher[orderDecorate] = append(m.decorateCypher[orderDecorate], d)
	return m
}

func (m *cypher) Where(con Condition) Cypher {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.conditions[m.this.getName()] = append(m.conditions[m.this.getName()], con)
	return m
}

func (m *cypher) LimitOffset(lim, os int64) Cypher {
	m.lock.Lock()
	defer m.lock.Unlock()
	d := &limitCypher{
		limit:  lim,
		offset: os,
	}
	m.decorateCypher[limitDecorate] = append(m.decorateCypher[limitDecorate], d)
	return m
}

func (m *cypher) Exec(ctx context.Context, connect connector.Connect) (*params_container.Result, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	cypherInput := m.parse()
	return connect.Exec(ctx, cypherInput, m.keys, m.accessMode)
}
func (m *cypher) TxExec(ctx context.Context, connect connector.Connect) (*params_container.Result, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	cypherInput := m.parse()
	return connect.TxExec(ctx, cypherInput, m.keys, m.accessMode)
}

func (m *cypher) parse() *params_container.CypherInput {
	var header = m.header
	var cypher = ""
	// 组建查询动作
	var respMap = make(map[string]interface{})
	var setters = make(map[string]*setter)
	var prwrite []PrimaryCypher
	for {
		if nil == header {
			break
		}
		if CypherMatch == header.action() || CypherDetachDelete == header.action() || CypherDelete == header.action() {
			cypherInput := header.cypher()
			cypher = fmt.Sprintf("%s %s", cypher, cypherInput.Cypher)
			for k, v := range cypherInput.Params {
				respMap[k] = v
			}
			//这块可以改为match create merge set这样的顺序，但是目前先这样，效率好像高一点
		} else if CypherUpdate == header.action() && !header.merge() {
			cypherInput := header.cypher()
			cypher = fmt.Sprintf("%s %s", cypher, cypherInput.Cypher)
			for k, v := range cypherInput.Params {
				respMap[k] = v
			}
			if nil != header.getProperty() && 0 < len(header.getProperty()) {
				setters[header.getName()] = &setter{header.getProperty()}
			}
		} else {
			prwrite = append(prwrite, header)
		}
		header = header.next()
	}

	// condition
	if nil != m.conditions && 0 < len(m.conditions) {
		var conditionStr = make([]string, 0, len(m.conditions))
		for name := range m.conditions {
			cons := m.conditions[name]

			for index := range cons {
				cy := cons[index].cypher(name)
				conditionStr = append(conditionStr, cy.Cypher)
				if cy.Params != nil && 0 < len(cy.Params) {
					for k, v := range cy.Params {
						respMap[k] = v
					}
				}
			}
		}
		cypher = fmt.Sprintf(" %s where %s ", cypher, strings.Join(conditionStr, " and "))
	}

	if nil != prwrite && 0 < len(prwrite) {
		for index := range prwrite {
			cypherInput := prwrite[index].cypher()
			cypher = fmt.Sprintf("%s %s", cypher, cypherInput.Cypher)
			for k, v := range cypherInput.Params {
				respMap[k] = v
			}
		}
	}
	if nil != setters && 0 < len(setters) {
		for name, set := range setters {
			setCypInput := set.cypher(name)
			for k, v := range setCypInput.Params {
				respMap[k] = v
			}
			cypher = fmt.Sprintf("%s %s", cypher, setCypInput.Cypher)
		}
	}
	if nil != m.decorateCypher {

		if nil != m.decorateCypher[returnDecorate] && 0 < len(m.decorateCypher[returnDecorate]) {
			returns := m.decorateCypher[returnDecorate]
			var returnStrs = make([]string, 0, len(returns))
			for index := range returns {
				ret := returns[index]
				returnCypherInput := ret.cypher()
				for k, v := range returnCypherInput.Params {
					respMap[k] = v
				}
				returnStrs = append(returnStrs, returnCypherInput.Cypher)
			}
			cypher = fmt.Sprintf("%s return %s ", cypher, strings.Join(returnStrs, ","))
		}
		if nil != m.decorateCypher[orderDecorate] && 0 < len(m.decorateCypher[orderDecorate]) {
			orders := m.decorateCypher[orderDecorate]
			var ordersStrs = make([]string, 0, len(orders))
			for index := range orders {
				orderCypherInput := orders[index].cypher()
				for k, v := range orderCypherInput.Params {
					respMap[k] = v
				}
				ordersStrs = append(ordersStrs, orderCypherInput.Cypher)
			}

			cypher = fmt.Sprintf("%s order by %s", cypher, strings.Join(ordersStrs, " , "))
		}
		if nil != m.decorateCypher[limitDecorate] && 0 < len(m.decorateCypher[limitDecorate]) {
			limit := m.decorateCypher[limitDecorate][0]
			limitCypherInput := limit.cypher()
			cypher = fmt.Sprintf("%s %s", cypher, limitCypherInput.Cypher)
		}
	}
	return &params_container.CypherInput{
		cypher,
		respMap,
	}
}

func (m *cypher) Plan() *params_container.CypherInput {

	cypher := m.parse()
	return cypher
}

func cypherActionToAccessMode(action CypherAction) neo4j.AccessMode {
	switch action {
	case CypherMatch:
		return neo4j.AccessModeRead
	default:
		return neo4j.AccessModeWrite
	}
}
