package params_container

type Relation struct {
	Name     string
	Property map[string]interface{}
}

type ParseNodeResp struct {
	NodeName string
	KeyName  string
	Property map[string]interface{}
}

type CypherInput struct {
	Cypher string
	Params map[string]interface{}
}

type Node struct {
	Label    string
	KeyName  string
	Property map[string]interface{}
}

type ParseRelationResp struct {
	Type     string
	Property map[string]interface{}
}

type Result struct {
	Data map[string][]interface{}
}
