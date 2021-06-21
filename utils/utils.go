package utils

import (
	"LinshanYu/neo4j-go-cypher/params_container"
	"errors"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"math/rand"
	"reflect"
	"time"
)

func GetRandomString(l int) string {
	str := "abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomString1(l int) string {
	str := "abcdefghijklmnopqrstuvwxyz0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetSuffix() string {
	return GetRandomString(1) + GetRandomString1(5)
}

func NodeConvert(nodeInterface interface{}) (*params_container.Node, error) {
	//neo4j.Relationship{}
	var reNode = &params_container.Node{}
	if node, ok := nodeInterface.(neo4j.Node); ok {
		if nil != node.Labels && 0 < len(node.Labels) {
			reNode.Label = node.Labels[0]
		}
		reNode.Property = node.Props
		return reNode, nil
	}
	if node, ok := nodeInterface.(*neo4j.Node); ok {
		if nil != node.Labels && 0 < len(node.Labels) {
			reNode.Label = node.Labels[0]
		}
		reNode.Property = node.Props
		return reNode, nil
	}
	return nil, errors.New(fmt.Sprintf("err type = %s convert to *params_container.Node", reflect.TypeOf(nodeInterface).String()))
}

func NodesConvert(nodes interface{}) ([]*params_container.Node, error) {
	var result []*params_container.Node
	if ints, ok := nodes.([]interface{}); ok {
		for index := range ints {
			node, err := NodeConvert(ints[index])
			if nil != err {
				return result, err
			}
			result = append(result, node)
		}
		return result, nil
	}
	if nNodes, ok := nodes.([]neo4j.Node); ok {
		for index := range nNodes {
			var reNode = &params_container.Node{}
			node := nNodes[index]
			if nil != node.Labels && 0 < len(node.Labels) {
				reNode.Label = node.Labels[0]
			}
			reNode.Property = node.Props
			result = append(result, reNode)
		}
		return result, nil
	}
	if nNodes, ok := nodes.([]*neo4j.Node); ok {
		for index := range nNodes {
			var reNode = &params_container.Node{}
			node := nNodes[index]
			if nil != node.Labels && 0 < len(node.Labels) {
				reNode.Label = node.Labels[0]
			}
			reNode.Property = node.Props
			result = append(result, reNode)
		}
		return result, nil
	}
	return nil, errors.New(fmt.Sprintf("err type = %s convert to []*params_container.Node", reflect.TypeOf(nodes).String()))
}

func RetryWithDelay(retry int, sleep time.Duration, fn func() error) error {
	if retry < 1 {
		return errors.New(fmt.Sprintf("error retry : %d", retry))
	}
	var returnErr error
	var index = 1
	for {
		returnErr = fn()
		if index == retry || returnErr == nil {
			return returnErr
		}
		time.Sleep(sleep)
		index++
	}
	return returnErr
}
