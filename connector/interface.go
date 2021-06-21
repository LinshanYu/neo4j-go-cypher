package connector

import (
	"LinshanYu/neo4j-go-cypher/params_container"
	"context"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Connect interface {
	TxExec(ctx context.Context, cypherInput *params_container.CypherInput, keys []string, accessMode neo4j.AccessMode) (*params_container.Result, error)
	Exec(ctx context.Context, cypherInput *params_container.CypherInput, keys []string, accessMode neo4j.AccessMode) (*params_container.Result, error)
}
