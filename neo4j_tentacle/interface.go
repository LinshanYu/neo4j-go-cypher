package neo4j_tentacle

import (
	"LinshanYu/neo4j-go-cypher/connector"
	"LinshanYu/neo4j-go-cypher/params_container"
	"LinshanYu/neo4j-go-cypher/utils"
	"context"
	"strings"
	"sync"
)

type CypherType string

var RelationType CypherType = "relation"
var NodeType CypherType = "node"

//Deprecated
var PathType CypherType = "path"

type CypherAction string

//write :update create
var CypherCreate CypherAction = "create"

var CypherUpdate CypherAction = "update"

var CypherDelete CypherAction = "delete"

var CypherDetachDelete CypherAction = "detachDelete"

//read :match
var CypherMatch CypherAction = "match"

type CypherRun interface {
	Exec(ctx context.Context, connect connector.Connect) (*params_container.Result, error)
	TxExec(ctx context.Context, connect connector.Connect) (*params_container.Result, error)
	Plan() *params_container.CypherInput
}

/**

cypher 关键字优先级

match
where
delete
merge/create
set
return
order by
limit


*/

type Cypher interface {
	ReNew(name string, cypherType CypherType, action CypherAction) Cypher
	//TODO 后续可能需要调整, 会和where整合,目前无好的解决办法
	//Deprecated
	//merge时 condition作为merge的条件
	//并且 对于 Relation 中传入 fromNodeAlias, toNodeAlias  fromNode, toNode , relations []string,也会作为merge的判断条件， 而relationProperty是set内容，此处需要谨慎
	Merge(condition map[string]interface{}) Cypher
	Node(node *params_container.Node) Cypher
	Relation(fromNodeAlias, toNodeAlias string, relations []string, relationProperty map[string]interface{}, direction bool) Cypher
	Return(distinct bool) Cypher
	OrderBy(propertyName string, asc bool) Cypher
	// Where 无法影响 Merge 操作
	Where(con Condition) Cypher
	LimitOffset(lim, os int64) Cypher
	Exec(ctx context.Context, connect connector.Connect) (*params_container.Result, error)
	TxExec(ctx context.Context, connect connector.Connect) (*params_container.Result, error)
	Plan() *params_container.CypherInput
}

// or and not in eql inlabels relation
// or and not 属于最外层的关联condition， 可以彼此嵌套
// eql in inlabel relation 属于内层判断条件condition【】
type Condition interface {
	cypher(name string) *params_container.CypherInput
	MatchAll(conditions ...Condition) Condition
	Match(condition Condition) Condition
}

func NewCypher(cypherName string, cypherType CypherType, action CypherAction) Cypher {

	var primaryCypher PrimaryCypher
	switch cypherType {
	case RelationType:
		primaryCypher = &relationCypher{
			ca: action,
		}
	case NodeType:
		primaryCypher = &nodeCypher{
			ca: action,
		}
	}
	cypherName = strings.Trim(cypherName, " ")
	if "" == strings.Trim(cypherName, " ") {
		cypherName = utils.GetRandomString(1) + utils.GetRandomString1(5)
	}
	primaryCypher.name(cypherName)
	m := &cypher{
		lock:           new(sync.Mutex),
		accessMode:     cypherActionToAccessMode(action),
		this:           primaryCypher,
		header:         primaryCypher,
		decorateCypher: make(map[decorateName][]DecorateCypher),
		conditions:     make(map[string][]Condition),
		keys:           []string{},
	}
	return m
}

func And() Condition {
	return &andCondition{}
}

func Or() Condition {
	return &orCondition{}
}

func Not() Condition {
	return &notCondition{}
}

func In(key string, values []interface{}) Condition {
	return &inCondition{
		Key:    key,
		Values: values,
	}
}

func InLabel(label string) Condition {
	return &inLabel{
		Label: label,
	}
}

func Eql(key string, value interface{}) Condition {
	return &eqlCondition{
		Key:   key,
		Value: value,
	}
}

func InRelation(relations []string, relationProperty map[string]interface{}, direction bool, opts ...func(in *inRelation)) Condition {
	inRe := &inRelation{
		RelationNames: relations,
		Property:      relationProperty,
		Direction:     direction,
	}
	for _, opt := range opts {
		opt(inRe)
	}
	return inRe
}

//当该node别名没有在前面或后面的条件中，就会报错
func InRelationFromAlias(fromAlias string) func(in *inRelation) {
	return func(in *inRelation) {
		in.FromNodeAlias = fromAlias
	}
}

//当该node别名没有在前面或后面的条件中，就会报错
func InRelationToAlias(toAlias string) func(in *inRelation) {
	return func(in *inRelation) {
		in.ToNodeAlias = toAlias
	}
}

//当该re别名没有在前面或者后面的条件中，就会报错
func InRelationAlias(reAlias string) func(in *inRelation) {
	return func(in *inRelation) {
		in.RelationAlias = reAlias
	}
}

type PrimaryCypher interface {
	cypher() *params_container.CypherInput
	name(name string)
	getName() string
	setNext(cy PrimaryCypher)
	next() PrimaryCypher
	action() CypherAction
	setLabel(labels string)
	setProperty(property map[string]interface{})
	getProperty() map[string]interface{}
	setRelations(relations []string, direction bool)
	setRelationFrom(nodeName string)
	setRelationTo(nodeName string)
	setMerge(condition map[string]interface{})
	merge() bool
}

type DecorateCypher interface {
	cypher() *params_container.CypherInput
}
