package utils

import (
	"LinshanYu/neo4j-go-cypher/params_container"
	"fmt"
	"strings"
)

func ParseMapColon(property map[string]interface{}, alias string) ([]string, map[string]interface{}) {
	if nil != property && 0 < len(property) {
		properties := make([]string, 0, len(property))
		var respMap = make(map[string]interface{}, len(property))
		suf := GetSuffix()
		for key := range property {
			keyName := alias + key + suf
			properties = append(properties, fmt.Sprintf("%s:$%s", key, keyName))
			respMap[keyName] = property[key]
		}
		return properties, respMap
	} else {
		return []string{}, map[string]interface{}{}
	}
}

func ParseMapEqual(property map[string]interface{}, alias string) ([]string, map[string]interface{}) {
	if nil != property && 0 < len(property) {
		properties := make([]string, 0, len(property))
		var respMap = make(map[string]interface{}, len(property))
		suf := GetSuffix()
		for key := range property {
			keyName := alias + key + suf
			properties = append(properties, fmt.Sprintf("%s.%s=$%s", alias, key, keyName))
			respMap[keyName] = property[key]
		}
		return properties, respMap
	} else {
		return []string{}, map[string]interface{}{}
	}
}

func ParseNodeProperty(nodeName, label string, properties map[string]interface{}) *params_container.CypherInput {
	nodeCypher := ""
	var respMap = make(map[string]interface{})
	if !("" == nodeName && "" == label) {
		var nodeAlis = ""
		if "" != nodeName {
			nodeAlis = fmt.Sprintf("%s", nodeName)
		}
		var labelStr = ""
		if "" != label {
			labelStr = fmt.Sprintf(":%s", label)

		}
		var propertyStr = ""
		if nil != properties && 0 < len(properties) {
			pros, res := ParseMapColon(properties, nodeName)
			propertyStr = strings.Join(pros, ",")
			for k, v := range res {
				respMap[k] = v
			}
		}
		nodeCypher = fmt.Sprintf("(%s%s{%s})", nodeAlis, labelStr, propertyStr)
	}
	return &params_container.CypherInput{
		Cypher: nodeCypher,
		Params: respMap,
	}
}

func ParseRelationBody(fromNodeName, toNodeName, reName string, relations []string, property map[string]interface{}, direction bool) *params_container.CypherInput {
	var respMap = make(map[string]interface{})
	var cypher = fmt.Sprintf("(%s)", fromNodeName)
	var reCypher = ""
	reAlis := fmt.Sprintf("%s", reName)

	relationsStr := ""
	if nil != relations && 0 < len(relations) {
		relationsStr = fmt.Sprintf(":%s", strings.Join(relations, "|"))
	}
	var propertyStr = ""
	if nil != property && 0 < len(property) {
		pros, res := ParseMapColon(property, reName)
		propertyStr = strings.Join(pros, ",")
		for k, v := range res {
			respMap[k] = v
		}
	}
	reCypher = fmt.Sprintf("%s%s{%s}", reAlis, relationsStr, propertyStr)
	var dir = ""
	if direction {
		dir = ">"
	}
	cypher = fmt.Sprintf("%s-[%s]-%s", cypher, reCypher, dir)
	cypher = fmt.Sprintf("%s(%s)", cypher, toNodeName)

	return &params_container.CypherInput{Cypher: cypher, Params: respMap}
}
