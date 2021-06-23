# 目标：
### &ensp; &ensp; 旨在为neo4j提供简易的cypher操作

# neo4j-go-cypher 提供以下接口：
## &ensp; &ensp; 接口列表
| 接口名称  |      接口说明      |  备注 |     实现  |
|------------|:------------------------|:---------------|:---------------|
| connector.Connect|  提供数据库读写能力 | 具有事务能力 |connector.neo4jDbConn 提供重试能力 |
| neo4j-tentacle.CypherRun |    cypher具有预执行和执行的能力   |   Plan返回该cypher预执行语句和参数 | neo4j-tentacle.Cypher |
| neo4j-tentacle.Cypher | 提供组装+执行cypher的能力  |    该接口面向用户，入口：neo4j-tentacle.NewCypher(cypherName string, cypherType CypherType, action CypherAction) Cypher ；目前支持两类三种操作（node+relation，以及update，create，match三种操作） | neo4j-tentacle.cypher |
| neo4j-tentacle.PrimaryCypher |  将一个cypher语句分为两部分，主部分（primary）+修饰部分（decorate），主部分为操作，修饰部分为对操作结果进行修饰（例如对结果order by等） |  该接口内部实现，用户不感知；基于不同的操作类型做不同的实现   | neo4j-tentacle.nodeCypher neo4j-tentacle.relationCypher |
| neo4j-tentacle.DecorateCypher | 修饰部分  |   目前包括：delete, return, limit, order | neo4j-tentacle.deleteCypher neo4j-tentacle.limitCypher neo4j-tentacle.orderCypher neo4j-tentacle.returnCypher|


## &ensp; &ensp; func 介绍
### &ensp; &ensp; &ensp; &ensp; 1. 入口：neo4j_tentacle.NewCypher(cypherName string, cypherType CypherType, action CypherAction) Cypher :
| 参数  |   位置 |   类型     |  参数说明 |     备注 |
|------------|:----|:------------------------|:---------------|:---------------|
| cypherName | input | string | 本次操作对象别名 | |
| cypherType | input |CypherType | 本次操作对象类型 | 目前支持RelationType和NodeType，其他类型待扩展|
| action | input | CypherAction | 本次操作类型| 目前支持CypherCreate CypherUpdate CypherDelete CypherDetachDelete CypherMatch |
| Cypher | output |Cypher | Cypher接口，后续所有操作基于该接口 | |

### &ensp; &ensp; &ensp; &ensp; 2. Cypher.ReNew(name string, cypherType CypherType, action CypherAction) Cypher : 基于原Cypher添加新cypher对象操作
| 参数  |   位置 |   类型     |  参数说明 |     备注 |
|------------|:----|:------------------------|:---------------|:---------------|
| name | input | string | 本次操作对象别名 | |
| cypherType | input |CypherType | 本次操作对象类型 | 目前支持RelationType和NodeType，其他类型待扩展|
| action | input | CypherAction | 本次操作类型| 目前支持CypherCreate CypherUpdate CypherDelete CypherDetachDelete CypherMatch |
| Cypher | output |Cypher | 返回ReNew后Cypher | |

### &ensp; &ensp; &ensp; &ensp; 3. Cypher.Node(node *params_container.Node) Cypher : NodeType操作时，传入node对象
| 参数  |   位置 |   类型     |  参数说明 |     备注 |
|------------|:----|:------------------------|:---------------|:---------------|
| node | input | *params_container.Node | node指针 | |
| Label  | node | string | node标签 | |
| Property | node | map[string]interface{} | node属性 | ｜

### &ensp; &ensp; &ensp; &ensp; 4. Cypher.Relation(fromNodeAlias, toNodeAlias string, relations []string, relationProperty map[string]interface{}, direction bool) Cypher : RelationType操作时，传入relation属性对象
| 参数  |   位置 |   类型     |  参数说明 |     备注 |
|------------|:----|:------------------------|:---------------|:---------------|
| fromNodeAlias | input | string | 关系起始节点别名 | |
| toNodeAlias  | input | string | 关系结束节点别名 | |
| relations  | input | []string | 关系类型 | |
| relationProperty  | input | map[string]interface{} | 关系属性 | |
| direction | input | bool | 关系是否具有指向性 | ｜

### &ensp; &ensp; &ensp; &ensp; ***5. Cypher.Merge(condition map[string]interface{}) Cypher : 对于Create或Update，基于merge提供的condition去判断是否已经存在某个节点：mergeCreate：无则创建，有则不创建； mergeUpdate：有则更新，无则创建***
| 参数  |   位置 |   类型     |  参数说明 |     备注 |
|------------|:----|:------------------------|:---------------|:---------------|
| condition | input | map[string]interface{} | merge基于的条件 | |

### &ensp; &ensp; &ensp; &ensp; 6. Cypher.Where(con Condition) Cypher : 判断条件，支持match/update/delete/detachDelete; ***不支持merge动作***
| 参数  |   位置 |   类型     |  参数说明 |     备注 |
|------------|:----|:------------------------|:---------------|:---------------|
| con | input | Condition | 判断条件 | Condition包括And，Or，Not，In，Eql，InLabel，InRelation等 |

### &ensp; &ensp; &ensp; &ensp; 7. Cypher.Return(distinct bool) Cypher : 返回当前操作的对象
| 参数  |   位置 |   类型     |  参数说明 |     备注 |
|------------|:----|:------------------------|:---------------|:---------------|
| distinct | input | bool | 返回数据去重 |  |

### &ensp; &ensp; &ensp; &ensp; 8. Cypher.OrderBy(propertyName string, asc bool) Cypher : 基于当前操作对象property排序
| 参数  |   位置 |   类型     |  参数说明 |     备注 |
|------------|:----|:------------------------|:---------------|:---------------|
| propertyName | input | string | 排序属性名 |  |
| asc | input | bool | 是否升序 |  |

### &ensp; &ensp; &ensp; &ensp; 9. Cypher.LimitOffset(lim, os int64) :
| 参数  |   位置 |   类型     |  参数说明 |     备注 |
|------------|:----|:------------------------|:---------------|:---------------|
| lim | input | int64 | 返回数量最大值 | 小于1， 返回所有 |
| os | input | int64 | 偏移量 | 小于1，返回所有 |

### &ensp; &ensp; &ensp; &ensp; 10. Cypher.Exec(ctx context.Context, connect connector.Connect) (*params_container.Result, error): 无事务执行cypher
| 参数  |   位置 |   类型     |  参数说明 |     备注 |
|------------|:----|:------------------------|:---------------|:---------------|
| connect | input | connector.Connect | neo4j连接 |  |
| result | input | *params_container.Result | 返回结果 |  params_container.Result.Data为map[string][]interface{}, 该map的key为操作对象别名，value为操作结果集合 |

### &ensp; &ensp; &ensp; &ensp; 11. Cypher.TxExec(ctx context.Context, connect connector.Connect) (*params_container.Result, error): 事务执行cypher
| 参数  |   位置 |   类型     |  参数说明 |     备注 |
|------------|:----|:------------------------|:---------------|:---------------|
| connect | input | connector.Connect | neo4j连接 |  |
| result | input | *params_container.Result | 返回结果 |  params_container.Result.Data为map[string][]interface{}, 该map的key为操作对象别名，value为操作结果集合 |
