package neo4j_tentacle

import (
	"LinshanYu/neo4j-go-cypher/params_container"
	"LinshanYu/neo4j-go-cypher/utils"
	"fmt"
	"strings"
)

type defaultCypher struct {
}

//Deprecated
func (d *defaultCypher) cypher() *params_container.CypherInput { return nil }

//Deprecated
func (d *defaultCypher) name(name string) { return }

//Deprecated
func (d *defaultCypher) getName() string { return "DeprecatedName" }

//Deprecated
func (d *defaultCypher) setNext(cy PrimaryCypher) {}

//Deprecated
func (d *defaultCypher) next() PrimaryCypher { return nil }

//Deprecated
func (d *defaultCypher) action() CypherAction { return CypherAction("DeprecatedAction") }

//Deprecated
func (d *defaultCypher) setLabel(labels string) {}

//Deprecated
func (d *defaultCypher) setProperty(property map[string]interface{}) {}

//Deprecated
func (d *defaultCypher) getProperty() map[string]interface{} {
	return map[string]interface{}{}
}

//Deprecated
func (d *defaultCypher) setRelations(relations []string, direction bool) {}

//Deprecated
func (d *defaultCypher) setRelationFrom(nodeName string) {}

//Deprecated
func (d *defaultCypher) setRelationTo(nodeName string) {}

//Deprecated
func (d *defaultCypher) setMerge(condition map[string]interface{}) {
}

//Deprecated
func (d *defaultCypher) merge() bool {
	return false
}

type nodeCypher struct {
	defaultCypher
	noName         string
	label          string
	property       map[string]interface{}
	nextCy         PrimaryCypher
	ca             CypherAction
	mergeFlag      bool
	mergeCondition map[string]interface{}
}

func (re *nodeCypher) cypher() *params_container.CypherInput {
	switch re.ca {
	case CypherMatch:
		return re.match()
	case CypherCreate:
		return re.create()
	case CypherUpdate:
		return re.update()
	default:
		return re.delete()
	}
	return nil
}

func (re *nodeCypher) setMerge(condition map[string]interface{}) {
	re.mergeFlag = true
	re.mergeCondition = condition
}
func (re *nodeCypher) merge() bool {
	return re.mergeFlag
}

func (re *nodeCypher) name(name string) {
	re.noName = name
}
func (re *nodeCypher) getName() string {
	return re.noName
}
func (re *nodeCypher) setNext(cy PrimaryCypher) {
	re.nextCy = cy
}
func (re *nodeCypher) next() PrimaryCypher {
	return re.nextCy
}

func (re *nodeCypher) setLabel(label string) {
	re.label = label
}

func (re *nodeCypher) setProperty(property map[string]interface{}) {
	re.property = property
}
func (re *nodeCypher) getProperty() map[string]interface{} {
	return re.property
}
func (re *nodeCypher) action() CypherAction {
	return re.ca
}

type relationCypher struct {
	defaultCypher
	fromNodeName   string
	toNodeName     string
	reName         string
	relations      []string
	property       map[string]interface{}
	nextCy         PrimaryCypher
	direction      bool
	ca             CypherAction
	mergeFlag      bool
	mergeCondition map[string]interface{}
}

func (re *relationCypher) cypher() *params_container.CypherInput {
	switch re.ca {
	case CypherMatch:
		return re.match()
	case CypherCreate:
		return re.create()
	case CypherUpdate:
		return re.update()
	default:
		return re.delete()
	}
	return nil

}
func (re *relationCypher) name(name string) {
	re.reName = name
}
func (re *relationCypher) getName() string {
	return re.reName
}
func (re *relationCypher) setNext(cy PrimaryCypher) {
	re.nextCy = cy
}
func (re *relationCypher) next() PrimaryCypher {
	return re.nextCy
}
func (re *relationCypher) action() CypherAction {
	return re.ca
}
func (re *relationCypher) setMerge(condition map[string]interface{}) {
	re.mergeFlag = true
	re.mergeCondition = condition
}

func (re *relationCypher) merge() bool {
	return re.mergeFlag
}

func (re *relationCypher) setProperty(property map[string]interface{}) {
	re.property = property
}
func (re *relationCypher) getProperty() map[string]interface{} {
	return re.property
}
func (re *relationCypher) setRelations(relations []string, direction bool) {
	re.relations = relations
	re.direction = direction
}
func (re *relationCypher) setRelationFrom(nodeName string) {
	re.fromNodeName = nodeName
}
func (re *relationCypher) setRelationTo(nodeName string) {
	re.toNodeName = nodeName
}

type setter struct {
	property map[string]interface{}
}

func (s *setter) cypher(name string) *params_container.CypherInput {

	cypher := ""

	properties, res := utils.ParseMapEqual(s.property, name)

	for _, p := range properties {
		cypher = fmt.Sprintf("%s set %s ", cypher, p)
	}

	return &params_container.CypherInput{
		cypher,
		res,
	}
}

func (re *nodeCypher) match() *params_container.CypherInput {
	labelStr := ""
	if re.label != "" {
		labelStr = ":" + re.label
	}
	var properties []string
	var respMap map[string]interface{}
	if nil != re.property && 0 < len(re.property) {
		properties = make([]string, 0, len(re.property))
		suf := utils.GetSuffix()
		respMap = make(map[string]interface{}, len(re.property))
		for key := range re.property {
			keyName := re.noName + key + suf
			properties = append(properties, fmt.Sprintf("%s:$%s", key, keyName))
			respMap[keyName] = re.property[key]
		}
	}

	cypher := fmt.Sprintf("match (%s%s{%s})", re.noName, labelStr, strings.Join(properties, ","))

	return &params_container.CypherInput{
		Cypher: cypher,
		Params: respMap,
	}
}

// merge create 无则创建基于set条件的数据 有则不做修改
// create 无条件创建
func (re *nodeCypher) create() *params_container.CypherInput {
	var respMap = make(map[string]interface{})
	var cypher = ""
	if re.mergeFlag {
		propertyCondition, condMap := utils.ParseMapColon(re.mergeCondition, re.noName)
		conditionStr := ""
		if nil != propertyCondition && 0 < len(propertyCondition) {
			conditionStr = strings.Join(propertyCondition, ",")
			for k, v := range condMap {
				respMap[k] = v
			}
		}
		cypher = fmt.Sprintf("merge (%s:%s{%s}) on create ", re.noName, re.label, conditionStr)

		propertySet, setMap := utils.ParseMapEqual(re.property, re.noName)
		if nil != propertySet && 0 < len(propertySet) {
			for _, v := range propertySet {
				cypher = fmt.Sprintf("%s set %s ", cypher, v)
			}
			for k, v := range setMap {
				respMap[k] = v
			}
		}
	} else {
		properties, res := utils.ParseMapColon(re.property, re.noName)
		for k, v := range res {
			respMap[k] = v
		}
		cypher = fmt.Sprintf("create (%s:%s{%s})", re.noName, re.label, strings.Join(properties, ","))
	}

	return &params_container.CypherInput{
		Cypher: cypher,
		Params: respMap,
	}
}

func (re *nodeCypher) update() *params_container.CypherInput {
	var respMap = make(map[string]interface{})
	var cypher = ""
	if re.mergeFlag {
		propertyCondition, condMap := utils.ParseMapColon(re.mergeCondition, re.noName)
		conditionStr := ""
		if nil != propertyCondition && 0 < len(propertyCondition) {
			conditionStr = strings.Join(propertyCondition, ",")
			for k, v := range condMap {
				respMap[k] = v
			}
		}
		cypher = fmt.Sprintf("merge (%s:%s{%s}) on match ", re.noName, re.label, conditionStr)

		propertySet, setMap := utils.ParseMapEqual(re.property, re.noName)
		if nil != propertySet && 0 < len(propertySet) {
			for _, v := range propertySet {
				cypher = fmt.Sprintf("%s set %s ", cypher, v)
			}
			for k, v := range setMap {
				respMap[k] = v
			}
		}
	} else {
		return &params_container.CypherInput{
			Cypher: fmt.Sprintf("match (%s:%s)", re.noName, re.label),
			Params: map[string]interface{}{},
		}
	}

	return &params_container.CypherInput{
		Cypher: cypher,
		Params: respMap,
	}
}

func (re *nodeCypher) delete() *params_container.CypherInput {

	ci := re.match()
	detach := ""
	if CypherDetachDelete == re.ca {
		detach = "detach"
	}
	ci.Cypher = fmt.Sprintf("%s %s delete %s", ci.Cypher, detach, re.noName)
	return ci
}

func (re *relationCypher) match() *params_container.CypherInput {
	result := utils.ParseRelationBody(re.fromNodeName, re.toNodeName, re.reName, re.relations, re.property, re.direction)
	result.Cypher = fmt.Sprintf("match %s", result.Cypher)
	return result
}

func (re *relationCypher) create() *params_container.CypherInput {
	if re.mergeFlag {
		result := utils.ParseRelationBody(re.fromNodeName, re.toNodeName, re.reName, re.relations, re.mergeCondition, re.direction)
		result.Cypher = fmt.Sprintf("merge %s", result.Cypher)
		pros, res := utils.ParseMapEqual(re.property, re.reName)
		if nil != pros && 0 < len(pros) {
			for _, v := range pros {
				result.Cypher = fmt.Sprintf("%s set %s ", result.Cypher, v)
			}
			for k, v := range res {
				result.Params[k] = v
			}
		}

		return result
	} else {
		result := utils.ParseRelationBody(re.fromNodeName, re.toNodeName, re.reName, re.relations, re.property, re.direction)
		result.Cypher = fmt.Sprintf("create %s", result.Cypher)
		return result
	}
}

func (re *relationCypher) update() *params_container.CypherInput {
	result := utils.ParseRelationBody(re.fromNodeName, re.toNodeName, re.reName, re.relations, re.mergeCondition, re.direction)
	if re.mergeFlag {
		result.Cypher = fmt.Sprintf("merge %s", result.Cypher)
		pros, res := utils.ParseMapEqual(re.property, re.reName)
		if nil != pros && 0 < len(pros) {
			for _, v := range pros {
				result.Cypher = fmt.Sprintf("%s set %s ", result.Cypher, v)
			}
			for k, v := range res {
				result.Params[k] = v
			}
		}

		return result
	} else {
		return &params_container.CypherInput{
			Cypher: fmt.Sprintf("match %s", result.Cypher),
			Params: result.Params,
		}
	}
}

func (re *relationCypher) delete() *params_container.CypherInput {
	ci := re.match()
	detach := ""
	if CypherDetachDelete == re.ca {
		detach = "detach"
	}
	ci.Cypher = fmt.Sprintf("%s %s delete %s", ci.Cypher, detach, re.reName)
	return ci
}
