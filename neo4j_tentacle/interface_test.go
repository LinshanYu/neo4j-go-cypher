package neo4j_tentacle

import (
	"LinshanYu/neo4j-go-cypher/connector"
	"LinshanYu/neo4j-go-cypher/params_container"
	"LinshanYu/neo4j-go-cypher/utils"
	"context"
	"testing"
)

var test_connect connector.Connect

func init() {
	conf := connector.NewConf("localhost", "7687")
	c, err := conf.GetConnt()
	if nil != err {
		panic(err)
	}
	test_connect = c
}

// Test list
// NewCypher
// Cypher.Plan
func Test_Match_NodeType(t *testing.T) {
	p := NewCypher("a", NodeType, CypherMatch).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
func Test_Create_NodeType(t *testing.T) {
	p := NewCypher("a", NodeType, CypherCreate).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
func Test_Update_NodeType(t *testing.T) {
	p := NewCypher("a", NodeType, CypherUpdate).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Merge
func Test_Create_Merge_NodeType(t *testing.T) {
	p := NewCypher("a", NodeType, CypherCreate).Merge(map[string]interface{}{"id": "123"}).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Merge
func Test_Update_Merge_NodeType(t *testing.T) {
	p := NewCypher("a", NodeType, CypherUpdate).Merge(map[string]interface{}{"id": "123"}).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Node
func Test_Create_NodeType_WithProperty(t *testing.T) {
	p := NewCypher("a", NodeType, CypherCreate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": "yls-test" + utils.GetSuffix()},
	}).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Node
func Test_Update_NodeType_WithProperty(t *testing.T) {
	p := NewCypher("a", NodeType, CypherUpdate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": "yls-test" + utils.GetSuffix()},
	}).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Merge
func Test_Create_Merge_NodeType_WithProperty(t *testing.T) {
	p := NewCypher("a", NodeType, CypherCreate).Merge(map[string]interface{}{"id": "123"}).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": "yls-test" + utils.GetSuffix()},
	}).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Merge
func Test_Update_Merge_NodeType_WithProperty(t *testing.T) {
	p := NewCypher("a", NodeType, CypherUpdate).Merge(map[string]interface{}{"id": "123"}).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": "yls-test" + utils.GetSuffix()},
	}).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Relation
func Test_Create_RelationType_WithProperty(t *testing.T) {
	from := &params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": "yls-test" + utils.GetSuffix()},
	}
	to := &params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": "yls-test" + utils.GetSuffix()},
	}
	p := NewCypher("from", NodeType, CypherMatch).Node(from).ReNew("to", NodeType, CypherMatch).Node(to).
		ReNew("a", RelationType, CypherCreate).Relation(
		"from", "to", []string{"SYSTEM"}, map[string]interface{}{"id": "123"}, true).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// Cypher.Plan
// Cypher.Relation
func Test_Update_RelationType_WithProperty(t *testing.T) {
	from := &params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": "yls-test" + utils.GetSuffix()},
	}
	to := &params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": "yls-test" + utils.GetSuffix()},
	}
	p := NewCypher("from", NodeType, CypherMatch).Node(from).ReNew("to", NodeType, CypherMatch).Node(to).
		ReNew("a", RelationType, CypherUpdate).Relation(
		"from", "to", []string{"SYSTEM"}, map[string]interface{}{"id": "123"}, true).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Relation
// Cypher.Merge
func Test_Create_Merge_RelationType_WithProperty(t *testing.T) {
	from := &params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": "yls-test" + utils.GetSuffix()},
	}
	to := &params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": "yls-test" + utils.GetSuffix()},
	}
	p := NewCypher("from", NodeType, CypherMatch).Node(from).ReNew("to", NodeType, CypherMatch).Node(to).
		ReNew("a", RelationType, CypherCreate).Relation(
		"from", "to", []string{"SYSTEM"}, map[string]interface{}{"id": "123"}, true).Merge(map[string]interface{}{"id": "123"}).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Relation
// Cypher.Merge
func Test_Update_Merge_RelationType_WithProperty(t *testing.T) {
	from := &params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": "yls-test" + utils.GetSuffix()},
	}
	to := &params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": "yls-test" + utils.GetSuffix()},
	}
	p := NewCypher("from", NodeType, CypherMatch).Node(from).ReNew("to", NodeType, CypherMatch).Node(to).
		ReNew("a", RelationType, CypherUpdate).Relation(
		"from", "to", []string{"SYSTEM"}, map[string]interface{}{"id": "123"}, true).Merge(map[string]interface{}{"id": "123"}).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
func Test_Match_Delete_NodeType(t *testing.T) {
	p := NewCypher("a", NodeType, CypherDelete).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
func Test_Match_Delete_Detach_NodeType(t *testing.T) {
	p := NewCypher("a", NodeType, CypherDetachDelete).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
func Test_Match_Delete_RelationType(t *testing.T) {
	p := NewCypher("a", RelationType, CypherDelete).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
func Test_Match_Delete_Detach_RelationType(t *testing.T) {
	p := NewCypher("a", RelationType, CypherDetachDelete).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Exec
func Test_Match_Delete_Detach_Exec_NodeType(t *testing.T) {
	result, errExec := NewCypher("a", RelationType, CypherDetachDelete).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"id": "yls-test"},
	}).Exec(context.Background(), test_connect)
	if nil != errExec {
		t.Fatal(errExec)
	}
	t.Log(result)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Return
func Test_Match_Return_NodeType(t *testing.T) {
	p := NewCypher("a", NodeType, CypherMatch).Return(false).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Return
func Test_Match_Return_Distinct_NodeType(t *testing.T) {
	p := NewCypher("a", NodeType, CypherMatch).Return(true).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Return
func Test_Match_Return_RelationType(t *testing.T) {
	p := NewCypher("a", RelationType, CypherMatch).Return(false).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Return
func Test_Match_Return_Distinct_RelationType(t *testing.T) {
	p := NewCypher("a", RelationType, CypherMatch).Return(true).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Return
// Cypher.OrderBy
func Test_Match_Return_Distinct_OrderBy_NodeType(t *testing.T) {
	p := NewCypher("a", NodeType, CypherMatch).Return(true).OrderBy("id", true).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Return
// Cypher.OrderBy
func Test_Match_Return_Distinct_OrderBy_RelationType(t *testing.T) {
	p := NewCypher("a", RelationType, CypherMatch).Return(true).OrderBy("id", true).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Return
// Cypher.Where
func Test_Match_Return_Distinct_Where_AndCondition_NodeType(t *testing.T) {

	or := Or().MatchAll(In("id", []interface{}{"yls-test123"}),
		InLabel("tag"),
		Eql("name", "zhangsan"),
		InRelation([]string{"SYSTEM"}, map[string]interface{}{"id": "345"}, true),
	)
	not := Not().MatchAll(
		In("id", []interface{}{"yls-test123"}),
		InLabel("tag"),
		Eql("name", "zhangsan"),
		InRelation([]string{"SYSTEM"}, map[string]interface{}{"id": "dshjak"}, true))
	and := And().MatchAll(or, not)

	p := NewCypher("a", NodeType, CypherMatch).Return(true).
		Where(and).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Return
// Cypher.Where
func Test_Match_Return_Distinct_Where_OrCondition_NodeType(t *testing.T) {

	and := And().MatchAll(In("id", []interface{}{"yls-test"}),
		InLabel("tag"),
		Eql("name", "zhangsan"),
		InRelation([]string{"SYSTEM"}, map[string]interface{}{"id": "dsajkdha"}, true),
	)
	not := Not().MatchAll(
		InLabel("tag"),
		Eql("name", "zhangsan"),
		InRelation([]string{"SYSTEM"}, map[string]interface{}{"id": "345"}, true))
	or := Or().MatchAll(and, not)

	p := NewCypher("a", NodeType, CypherMatch).Return(true).
		Where(or).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Return
// Cypher.Where
func Test_Match_Return_Distinct_Where_NotCondition_NodeType(t *testing.T) {

	and := And().MatchAll(In("id", []interface{}{"yls-test"}),
		InLabel("tag"),
		Eql("name", "zhangsan"),
		InRelation([]string{"SYSTEM"}, map[string]interface{}{"id": "dsajkdha"}, true),
	)
	or := Or().MatchAll(
		InLabel("tag"),
		Eql("name", "zhangsan"),
		InRelation([]string{"SYSTEM"}, map[string]interface{}{"id": "345"}, true))
	not := Not().MatchAll(and, or)

	p := NewCypher("a", NodeType, CypherMatch).Return(true).
		Where(not).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Return
// Cypher.LimitOffset
// error LimitOrOffset will match all
func Test_Match_Return_Distinct_LimitOffsetErrLimit_NodeType(t *testing.T) {

	p := NewCypher("a", NodeType, CypherMatch).Return(true).
		LimitOffset(-1, 0).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Return
// Cypher.LimitOffset
// error LimitOrOffset will match all
func Test_Match_Return_Distinct_LimitOffsetErrOffset_NodeType(t *testing.T) {

	p := NewCypher("a", NodeType, CypherMatch).Return(true).
		LimitOffset(1, -1).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Plan
// Cypher.Return
// Cypher.LimitOffset
func Test_Match_Return_Distinct_LimitOffset_NodeType(t *testing.T) {

	p := NewCypher("a", NodeType, CypherMatch).Return(true).
		LimitOffset(10, 0).Plan()
	t.Log(p.Cypher)
	t.Log(p.Params)
}

// Test list
// NewCypher
// Cypher.Return
// Cypher.LimitOffset
// Cypher.Exec
func Test_Match_Return_Distinct_LimitOffset_ExecWrite_NodeType(t *testing.T) {

	result, err := NewCypher("a", NodeType, CypherMatch).Return(true).
		LimitOffset(10, 0).Exec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}

	if nil != result && nil != result.Data && 0 < len(result.Data) {
		data := result.Data["a"]
		if nil != data && 0 < len(data) {
			nodes, err := utils.NodesConvert(data)
			if nil != err {
				t.Fatal(err)
			}
			t.Log(len(nodes))
			t.Log(nodes[0].Label)
		}
	}
}

// Test list
// NewCypher
// Cypher.Return
// Cypher.Node
// Cypher.Exec
func Test_Create_ExecWrite_NodeType(t *testing.T) {

	result, err := NewCypher("a", NodeType, CypherCreate).
		Node(&params_container.Node{
			Label:    "tag",
			KeyName:  "",
			Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
		}).Return(true).Exec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}

	if nil != result && nil != result.Data && 0 < len(result.Data) {
		data := result.Data["a"]
		if nil != data && 0 < len(data) {
			nodes, err := utils.NodesConvert(data)
			if nil != err {
				t.Fatal(err)
			}
			t.Log(len(nodes))
			t.Log(nodes[0].Label)
		}
	}
}

// Test list
// NewCypher
// Cypher.Return
// Cypher.Node
// Cypher.Exec
func Test_Create_Merge_ExecWrite_NodeType(t *testing.T) {

	result, err := NewCypher("a", NodeType, CypherCreate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).Merge(map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()}).Return(true).
		Exec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}

	if nil != result && nil != result.Data && 0 < len(result.Data) {
		data := result.Data["a"]
		if nil != data && 0 < len(data) {
			nodes, err := utils.NodesConvert(data)
			if nil != err {
				t.Fatal(err)
			}
			t.Log(len(nodes))
			t.Log(nodes[0].Label)
		}
	}
}

// Test list
// NewCypher
// Cypher.Return
// Cypher.Node
// Cypher.Exec
func Test_Update_ExecWrite_NodeType(t *testing.T) {

	result, err := NewCypher("a", NodeType, CypherUpdate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).Return(true).
		Exec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}

	if nil != result && nil != result.Data && 0 < len(result.Data) {
		data := result.Data["a"]
		if nil != data && 0 < len(data) {
			nodes, err := utils.NodesConvert(data)
			if nil != err {
				t.Fatal(err)
			}
			t.Log(len(nodes))
			t.Log(nodes[0].Label)
		}
	}
}

// Test list
// NewCypher
// Cypher.Return
// Cypher.Node
// Cypher.Where
// Cypher.Exec
func Test_Update_Where_ExecWrite_NodeType(t *testing.T) {

	result, err := NewCypher("a", NodeType, CypherUpdate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).Where(And().MatchAll(
		Eql("name", "zhangsan"),
	)).Return(true).
		Exec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}

	if nil != result && nil != result.Data && 0 < len(result.Data) {
		data := result.Data["a"]
		if nil != data && 0 < len(data) {
			nodes, err := utils.NodesConvert(data)
			if nil != err {
				t.Fatal(err)
			}
			t.Log(len(nodes))
			t.Log(nodes[0].Label)
		}
	}
}

// Test list
// NewCypher
// Cypher.Return
// Cypher.LimitOffset
// Cypher.Exec
func Test_Update_Merge_ExecWrite_NodeType(t *testing.T) {

	result, err := NewCypher("a", NodeType, CypherUpdate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).Merge(map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()}).Return(true).
		Exec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}

	if nil != result && nil != result.Data && 0 < len(result.Data) {
		data := result.Data["a"]
		if nil != data && 0 < len(data) {
			nodes, err := utils.NodesConvert(data)
			if nil != err {
				t.Fatal(err)
			}
			t.Log(len(nodes))
			t.Log(nodes[0].Label)
		}
	}
}

// Test list
// NewCypher
// Cypher.Return
// Cypher.LimitOffset
// Cypher.TxExec
func Test_Match_Return_Distinct_LimitOffset_TxExecWrite_NodeType(t *testing.T) {

	result, err := NewCypher("a", NodeType, CypherMatch).Return(true).
		LimitOffset(10, 0).TxExec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}

	if nil != result && nil != result.Data && 0 < len(result.Data) {
		data := result.Data["a"]
		if nil != data && 0 < len(data) {
			nodes, err := utils.NodesConvert(data)
			if nil != err {
				t.Fatal(err)
			}
			t.Log(len(nodes))
			t.Log(nodes[0].Label)
		}
	}
}

// Test list
// NewCypher
// Cypher.Return
// Cypher.Node
// Cypher.TxExec
func Test_Create_TxExecWrite_NodeType(t *testing.T) {

	result, err := NewCypher("a", NodeType, CypherCreate).
		Node(&params_container.Node{
			Label:    "tag",
			KeyName:  "",
			Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
		}).Return(true).TxExec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}

	if nil != result && nil != result.Data && 0 < len(result.Data) {
		data := result.Data["a"]
		if nil != data && 0 < len(data) {
			nodes, err := utils.NodesConvert(data)
			if nil != err {
				t.Fatal(err)
			}
			t.Log(len(nodes))
			t.Log(nodes[0].Label)
		}
	}
}

// Test list
// NewCypher
// Cypher.Return
// Cypher.Node
// Cypher.TxExec
func Test_Create_Merge_TxExecWrite_NodeType(t *testing.T) {

	result, err := NewCypher("a", NodeType, CypherCreate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).Merge(map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()}).Return(true).
		TxExec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}

	if nil != result && nil != result.Data && 0 < len(result.Data) {
		data := result.Data["a"]
		if nil != data && 0 < len(data) {
			nodes, err := utils.NodesConvert(data)
			if nil != err {
				t.Fatal(err)
			}
			t.Log(len(nodes))
			t.Log(nodes[0].Label)
		}
	}
}

// Test list
// NewCypher
// Cypher.Return
// Cypher.Node
// Cypher.TxExec
func Test_Update_TxExecWrite_NodeType(t *testing.T) {

	result, err := NewCypher("a", NodeType, CypherUpdate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).Return(true).
		TxExec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}

	if nil != result && nil != result.Data && 0 < len(result.Data) {
		data := result.Data["a"]
		if nil != data && 0 < len(data) {
			nodes, err := utils.NodesConvert(data)
			if nil != err {
				t.Fatal(err)
			}
			t.Log(len(nodes))
			t.Log(nodes[0].Label)
		}
	}
}

// Test list
// NewCypher
// Cypher.Return
// Cypher.Node
// Cypher.Where
// Cypher.TxExec
func Test_Update_Where_TxExecWrite_NodeType(t *testing.T) {

	result, err := NewCypher("a", NodeType, CypherUpdate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).Where(And().MatchAll(
		Eql("name", "zhangsan"),
	)).Return(true).
		TxExec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}

	if nil != result && nil != result.Data && 0 < len(result.Data) {
		data := result.Data["a"]
		if nil != data && 0 < len(data) {
			nodes, err := utils.NodesConvert(data)
			if nil != err {
				t.Fatal(err)
			}
			t.Log(len(nodes))
			t.Log(nodes[0].Label)
		}
	}
}

// Test list
// NewCypher
// Cypher.Return
// Cypher.LimitOffset
// Cypher.TxExec
func Test_Update_Merge_TxExecWrite_NodeType(t *testing.T) {

	result, err := NewCypher("a", NodeType, CypherUpdate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).Merge(map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()}).Return(true).
		TxExec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}

	if nil != result && nil != result.Data && 0 < len(result.Data) {
		data := result.Data["a"]
		if nil != data && 0 < len(data) {
			nodes, err := utils.NodesConvert(data)
			if nil != err {
				t.Fatal(err)
			}
			t.Log(len(nodes))
			t.Log(nodes[0].Label)
		}
	}
}

// Test list
// NewCypher
// Cypher.Return
// Cypher.LimitOffset
// Cypher.Exec
func Test_Match_Return_Distinct_LimitOffset_ExecRead_NodeType(t *testing.T) {

	result, err := NewCypher("a", NodeType, CypherMatch).Return(true).
		LimitOffset(10, 0).Exec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}

	if nil != result && nil != result.Data && 0 < len(result.Data) {
		data := result.Data["a"]
		if nil != data && 0 < len(data) {
			nodes, err := utils.NodesConvert(data)
			if nil != err {
				t.Fatal(err)
			}
			t.Log(len(nodes))
			t.Log(nodes[0].Label)
		}
	}
}

// Test list
// NewCypher
// Cypher.Return
// Cypher.LimitOffset
// Cypher.TxExec
func Test_Match_Return_Distinct_LimitOffset_TxExecRead_NodeType(t *testing.T) {

	result, err := NewCypher("a", NodeType, CypherMatch).Return(true).
		LimitOffset(10, 0).TxExec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}

	if nil != result && nil != result.Data && 0 < len(result.Data) {
		data := result.Data["a"]
		if nil != data && 0 < len(data) {
			nodes, err := utils.NodesConvert(data)
			if nil != err {
				t.Fatal(err)
			}
			t.Log(len(nodes))
			t.Log(nodes[0].Label)
		}
	}
}

// Test list
// NewCypher
// Cypher.Node
// Cypher.ReNew
// Cypher.Where
// Cypher.TxExec
func Test_Renew_MatchAndUpdate_NodeType(t *testing.T) {

	_, err := NewCypher("a", NodeType, CypherMatch).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).ReNew("b", NodeType, CypherUpdate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).Where(And().Match(Eql("id", "yls-test"))).
		TxExec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}
}

// Test list
// NewCypher
// Cypher.Node
// Cypher.ReNew
// Cypher.Merge
// Cypher.TxExec
func Test_Renew_MatchAndMergeUpdate_NodeType(t *testing.T) {

	_, err := NewCypher("a", NodeType, CypherMatch).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).ReNew("b", NodeType, CypherUpdate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).Merge(map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()}).
		TxExec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}
}

// Test list
// NewCypher
// Cypher.Node
// Cypher.ReNew
// Cypher.TxExec
func Test_Renew_MatchAndCreate_NodeType(t *testing.T) {

	_, err := NewCypher("a", NodeType, CypherMatch).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).ReNew("b", NodeType, CypherCreate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).
		TxExec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}
}

// Test list
// NewCypher
// Cypher.Node
// Cypher.ReNew
// Cypher.Merge
// Cypher.TxExec
func Test_Renew_MatchAndMergeCreate_NodeType(t *testing.T) {

	_, err := NewCypher("a", NodeType, CypherMatch).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).ReNew("b", NodeType, CypherCreate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).Merge(map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()}).
		TxExec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}
}

// Test list
// NewCypher
// Cypher.Node
// Cypher.Where
// Cypher.Return
// Cypher.ReNew
// Cypher.TxExec
func Test_Renew_UpdateAndMatch_NodeType(t *testing.T) {

	_, err := NewCypher("a", NodeType, CypherUpdate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).Where(And().MatchAll(
		Eql("name", "zhangsan"),
		Eql("id", "yls-test"),
	)).Return(true).ReNew("b", NodeType, CypherMatch).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).
		TxExec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}
}

// Test list
// NewCypher
// Cypher.Node
// Cypher.Merge
// Cypher.Return
// Cypher.ReNew
// Cypher.TxExec
func Test_Renew_MergeUpdateAndMatch_NodeType(t *testing.T) {

	_, err := NewCypher("a", NodeType, CypherUpdate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix(), "id": utils.GetSuffix()},
	}).Merge(map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()}).Return(true).ReNew("b", NodeType, CypherMatch).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix(), "id": utils.GetSuffix()},
	}).
		TxExec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}
}

// Test list
// NewCypher
// Cypher.Node
// Cypher.Return
// Cypher.ReNew
// Cypher.TxExec
func Test_Renew_CreateAndMatch_NodeType(t *testing.T) {

	_, err := NewCypher("a", NodeType, CypherCreate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).Return(true).ReNew("b", NodeType, CypherMatch).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()},
	}).
		TxExec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}
}

// Test list
// NewCypher
// Cypher.Node
// Cypher.Merge
// Cypher.Return
// Cypher.ReNew
// Cypher.TxExec
func Test_Renew_MergeCreateAndMatch_NodeType(t *testing.T) {

	_, err := NewCypher("a", NodeType, CypherCreate).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix(), "id": utils.GetSuffix()},
	}).Merge(map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix()}).Return(true).ReNew("b", NodeType, CypherMatch).Node(&params_container.Node{
		Label:    "tag",
		KeyName:  "",
		Property: map[string]interface{}{"name": utils.GetSuffix() + utils.GetSuffix(), "id": utils.GetSuffix()},
	}).
		TxExec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}
}

//TODO 关于relationn的动作测试较少，重点关注merge relation
//以下测试relation
//match detachDelete delete update create mergeCreate
//reNew场景下亦执行以上场景

func Test_Match_Exec_RelationType(t *testing.T) {

	from := &params_container.Node{
		Label:    "pig",
		KeyName:  "",
		Property: map[string]interface{}{"id": utils.GetSuffix()},
	}

	to := &params_container.Node{
		Label:    "pig",
		KeyName:  "",
		Property: map[string]interface{}{"id": utils.GetSuffix()},
	}

	result, err := NewCypher("from", NodeType, CypherMatch).Node(from).ReNew("to", NodeType, CypherMatch).Node(to).
		ReNew("re", RelationType, CypherMatch).Relation("a", "b",
		[]string{"SYSTEM"},
		map[string]interface{}{"id": "123"},
		true,
	).Return(true).Exec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(result)
}

func Test_Delete_Exec_RelationType(t *testing.T) {

	from := &params_container.Node{
		Label:    "pig",
		KeyName:  "",
		Property: map[string]interface{}{"id": utils.GetSuffix()},
	}

	to := &params_container.Node{
		Label:    "pig",
		KeyName:  "",
		Property: map[string]interface{}{"id": utils.GetSuffix()},
	}

	result, err := NewCypher("from", NodeType, CypherMatch).Node(from).ReNew("to", NodeType, CypherMatch).Node(to).
		ReNew("re", RelationType, CypherDelete).Relation("a", "b",
		[]string{"SYSTEM"},
		map[string]interface{}{"id": "123"},
		true,
	).Return(true).Exec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(result)
}

func Test_DetachDelete_Exec_RelationType(t *testing.T) {

	from := &params_container.Node{
		Label:    "pig",
		KeyName:  "",
		Property: map[string]interface{}{"id": utils.GetSuffix()},
	}

	to := &params_container.Node{
		Label:    "pig",
		KeyName:  "",
		Property: map[string]interface{}{"id": utils.GetSuffix()},
	}

	result, err := NewCypher("from", NodeType, CypherMatch).Node(from).ReNew("to", NodeType, CypherMatch).Node(to).
		ReNew("re", RelationType, CypherDetachDelete).Relation("a", "b",
		[]string{"SYSTEM"},
		map[string]interface{}{"id": "123"},
		true,
	).Return(true).Exec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(result)

}
func Test_Update_Exec_RelationType(t *testing.T) {

	from := &params_container.Node{
		Label:    "pig",
		KeyName:  "",
		Property: map[string]interface{}{"id": utils.GetSuffix()},
	}

	to := &params_container.Node{
		Label:    "pig",
		KeyName:  "",
		Property: map[string]interface{}{"id": utils.GetSuffix()},
	}

	result, err := NewCypher("from", NodeType, CypherMatch).Node(from).ReNew("to", NodeType, CypherMatch).Node(to).
		ReNew("re", RelationType, CypherUpdate).Relation("a", "b",
		[]string{"SYSTEM"},
		map[string]interface{}{"id": "123"},
		true,
	).Return(true).Exec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(result)

}
func Test_MergeUpdate_Exec_RelationType(t *testing.T) {
	from := &params_container.Node{
		Label:    "pig",
		KeyName:  "",
		Property: map[string]interface{}{"id": utils.GetSuffix()},
	}

	to := &params_container.Node{
		Label:    "pig",
		KeyName:  "",
		Property: map[string]interface{}{"id": utils.GetSuffix()},
	}

	result, err := NewCypher("from", NodeType, CypherMatch).Node(from).ReNew("to", NodeType, CypherMatch).Node(to).
		ReNew("re", RelationType, CypherUpdate).Relation("a", "b",
		[]string{"SYSTEM"},
		map[string]interface{}{"id": "123"},
		true,
	).Merge(map[string]interface{}{"ylstest": "123dsahjkdsa"}).Return(true).Exec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(result)
}

func Test_Create_Exec_RelationType(t *testing.T) {

	from := &params_container.Node{
		Label:    "pig",
		KeyName:  "",
		Property: map[string]interface{}{"id": utils.GetSuffix()},
	}

	to := &params_container.Node{
		Label:    "pig",
		KeyName:  "",
		Property: map[string]interface{}{"id": utils.GetSuffix()},
	}

	result, err := NewCypher("from", NodeType, CypherMatch).Node(from).ReNew("to", NodeType, CypherMatch).Node(to).
		ReNew("re", RelationType, CypherCreate).Relation("a", "b",
		[]string{"SYSTEM"},
		map[string]interface{}{"id": "123"},
		true,
	).Return(true).Exec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(result)

}

func Test_MergeCreate_Exec_RelationType(t *testing.T) {
	from := &params_container.Node{
		Label:    "pig",
		KeyName:  "",
		Property: map[string]interface{}{"id": utils.GetSuffix()},
	}

	to := &params_container.Node{
		Label:    "pig",
		KeyName:  "",
		Property: map[string]interface{}{"id": utils.GetSuffix()},
	}

	result, err := NewCypher("from", NodeType, CypherMatch).Node(from).ReNew("to", NodeType, CypherMatch).Node(to).
		ReNew("re", RelationType, CypherCreate).Relation("a", "b",
		[]string{"SYSTEM"},
		map[string]interface{}{"id": "123"},
		true,
	).Merge(map[string]interface{}{"ylstest": "123dsahjkdsa"}).Return(true).Exec(context.Background(), test_connect)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(result)
}
