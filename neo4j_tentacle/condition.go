package neo4j_tentacle

import (
	"LinshanYu/neo4j-go-cypher/params_container"
	"LinshanYu/neo4j-go-cypher/utils"
	"fmt"
	"strings"
)

type andCondition struct {
	conditions []Condition
}

func (and *andCondition) MatchAll(conditions ...Condition) Condition {
	and.conditions = append(and.conditions, conditions...)
	return and
}
func (and *andCondition) Match(condition Condition) Condition {
	and.conditions = append(and.conditions, condition)
	return and
}

func (and *andCondition) cypher(name string) *params_container.CypherInput {
	var respMap = make(map[string]interface{})
	var cypher = ""
	var condStr = make([]string, 0, len(and.conditions))
	for index := range and.conditions {
		cond := and.conditions[index]
		cypherInput := cond.cypher(name)
		for k, v := range cypherInput.Params {
			respMap[k] = v
		}
		condStr = append(condStr, fmt.Sprintf("%s", cypherInput.Cypher))
	}
	cypher = fmt.Sprintf("(%s)", strings.Join(condStr, " and "))
	return &params_container.CypherInput{
		Cypher: cypher,
		Params: respMap,
	}
}

type orCondition struct {
	conditions []Condition
}

func (and *orCondition) MatchAll(conditions ...Condition) Condition {
	and.conditions = append(and.conditions, conditions...)
	return and
}
func (and *orCondition) Match(condition Condition) Condition {
	and.conditions = append(and.conditions, condition)
	return and
}

func (and *orCondition) cypher(name string) *params_container.CypherInput {
	var respMap = make(map[string]interface{})
	var cypher = ""
	var condStr = make([]string, 0, len(and.conditions))
	for index := range and.conditions {
		cond := and.conditions[index]
		cypherInput := cond.cypher(name)
		for k, v := range cypherInput.Params {
			respMap[k] = v
		}
		condStr = append(condStr, fmt.Sprintf("%s", cypherInput.Cypher))
	}
	cypher = fmt.Sprintf("(%s)", strings.Join(condStr, " or "))
	return &params_container.CypherInput{
		Cypher: cypher,
		Params: respMap,
	}
}

type notCondition struct {
	conditions []Condition
}

func (e *notCondition) cypher(name string) *params_container.CypherInput {

	var respMap = make(map[string]interface{})
	var cys = make([]string, 0, len(e.conditions))
	for index := range e.conditions {
		con := e.conditions[index]
		cypInput := con.cypher(name)
		cys = append(cys, cypInput.Cypher)
		for k, v := range cypInput.Params {
			respMap[k] = v
		}
	}
	var cypher = fmt.Sprintf("(not (%s))", strings.Join(cys, " and "))

	return &params_container.CypherInput{
		cypher,
		respMap,
	}
}

func (e *notCondition) MatchAll(conditions ...Condition) Condition {
	e.conditions = append(e.conditions, conditions...)
	return e
}

func (e *notCondition) Match(condition Condition) Condition {
	e.conditions = append(e.conditions, condition)
	return e
}

type defaultCondition struct {
}

//Deprecated
func (e *defaultCondition) cypher(name string) *params_container.CypherInput {
	return nil
}

//Deprecated
func (e *defaultCondition) MatchAll(conditions ...Condition) Condition {
	return e
}

//Deprecated
func (e *defaultCondition) Match(condition Condition) Condition {
	return e
}

type inCondition struct {
	defaultCondition
	Key    string
	Values []interface{}
}

func (e *inCondition) cypher(name string) *params_container.CypherInput {
	//TODO	还没摸透neo4j-go-driver传入数组的解决办法 只能先打散再传入

	var respMap = make(map[string]interface{}, len(e.Values))
	var valueStrs = make([]string, 0, len(e.Values))

	for index := range e.Values {
		value := e.Values[index]
		suf := utils.GetSuffix()
		valueStrs = append(valueStrs, fmt.Sprintf("$%s%s%s", name, e.Key, suf))
		respMap[name+e.Key+suf] = value
	}

	cypher := fmt.Sprintf("(%s.%s in [%s])", name, e.Key, strings.Join(valueStrs, ","))

	return &params_container.CypherInput{
		cypher,
		respMap,
	}
}

type inRelation struct {
	defaultCondition
	//Deprecated
	//当该node别名没有在前面的条件或者查询或者写入出现过时，就会报错，以下两点同理
	FromNodeAlias string
	//Deprecated
	ToNodeAlias string
	//Deprecated
	RelationAlias string
	RelationNames []string
	Property      map[string]interface{}
	Direction     bool
}

func (e *inRelation) cypher(name string) *params_container.CypherInput {
	relations := ""
	if nil != e.RelationNames && 0 < len(e.RelationNames) {
		relations = ":" + strings.Join(e.RelationNames, "|")
	}
	var propertyStr []string
	var respMap map[string]interface{}
	if nil != e.Property && 0 < len(e.Property) {
		propertyStr = make([]string, 0, len(e.Property))
		suf := utils.GetSuffix()
		respMap = make(map[string]interface{}, len(e.Property))
		for k, v := range e.Property {
			keyName := name + k + suf
			propertyStr = append(propertyStr, fmt.Sprintf("%s:$%s", k, keyName))
			respMap[keyName] = v
		}
	}
	var dir = ""
	if e.Direction {
		dir = ">"
	}
	cypher := fmt.Sprintf("((%s)-[%s%s{%s}]-%s(%s))", e.FromNodeAlias, e.RelationAlias, relations, strings.Join(propertyStr, ","), dir, e.ToNodeAlias)

	return &params_container.CypherInput{cypher, respMap}
}

type eqlCondition struct {
	defaultCondition
	Key   string
	Value interface{}
}

func (e *eqlCondition) cypher(name string) *params_container.CypherInput {
	suf := utils.GetSuffix()
	cypher := fmt.Sprintf("(%s.%s=$%s%s%s)", name, e.Key, name, e.Key, suf)

	return &params_container.CypherInput{
		cypher,
		map[string]interface{}{name + e.Key + suf: e.Value},
	}
}

type inLabel struct {
	defaultCondition
	Label string
}

func (e *inLabel) cypher(name string) *params_container.CypherInput {

	cypher := fmt.Sprintf("(\"%s\" in labels(%s))", e.Label, name)

	return &params_container.CypherInput{
		cypher,
		map[string]interface{}{},
	}
}
