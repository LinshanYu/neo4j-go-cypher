package neo4j_tentacle

import (
	"LinshanYu/neo4j-go-cypher/params_container"
	"fmt"
)

type limitCypher struct {
	limit  int64
	offset int64
}

func (d *limitCypher) cypher() *params_container.CypherInput {
	if d.offset < 0 || d.limit < 1 {
		return &params_container.CypherInput{}
	}
	return &params_container.CypherInput{
		Cypher: fmt.Sprintf(" skip %d limit %d ", d.offset, d.limit),
	}
}

type orderCypher struct {
	nodeName string
	order    string
	asc      bool
}

func (r *orderCypher) cypher() *params_container.CypherInput {

	asc := "desc"
	if r.asc {
		asc = "asc"
	}
	return &params_container.CypherInput{
		Cypher: fmt.Sprintf(" %s.%s %s", r.nodeName, r.order, asc),
		Params: nil,
	}
}

type returnCypher struct {
	nodeName string
	distinct bool
}

func (r *returnCypher) cypher() *params_container.CypherInput {

	if r.distinct {
		return &params_container.CypherInput{
			Cypher: fmt.Sprintf(" distinct( %s ) as %s ", r.nodeName, r.nodeName),
		}
	}
	return &params_container.CypherInput{
		Cypher: fmt.Sprintf(" %s ", r.nodeName),
	}
}
