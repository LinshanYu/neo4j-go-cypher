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

