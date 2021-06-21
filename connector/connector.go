package connector

import (
	"LinshanYu/neo4j-go-cypher/params_container"
	"LinshanYu/neo4j-go-cypher/utils"
	"context"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"math/rand"
	"strings"
	"time"
)

type Neo4jConf struct {
	Host         string
	Port         string
	Username     string
	Password     string
	Database     string
	ReadTimeout  int16
	WriteTimeout int16
	DbVersion    string
	Realm        string
}

func NewConf(host, port string, confs ...func(conf *Neo4jConf)) *Neo4jConf {
	config := &Neo4jConf{
		Host: host,
		Port: port,
	}
	for _, conf := range confs {
		conf(config)
	}
	return config
}

func (conf *Neo4jConf) GetConnt() (Connect, error) {
	var dbUri = "bolt+ssc://" + conf.Host + ":" + conf.Port
	if "4" == strings.Trim(strings.Split(conf.DbVersion, ".")[0], " ") {
		dbUri = "neo4j://" + conf.Host + ":" + conf.Port
	}
	dri, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth(conf.Username, conf.Password, conf.Realm))
	if nil != err {
		return nil, err
	}
	return &neo4jDbConn{
		database:     conf.Database,
		readTimeout:  time.Duration(conf.ReadTimeout) * time.Millisecond,
		writeTimeout: time.Duration(conf.WriteTimeout) * time.Millisecond,
		driver:       dri,
	}, nil
}

func WithUsernameAndPassword(username, password string) func(conf *Neo4jConf) {
	return func(conf *Neo4jConf) {
		conf.Username = username
		conf.Password = password
	}
}

func WithDatabase(database string) func(conf *Neo4jConf) {
	return func(conf *Neo4jConf) {
		conf.Database = database
	}
}

func WithReadTimeout(readTimeout int16) func(conf *Neo4jConf) {
	return func(conf *Neo4jConf) {
		conf.ReadTimeout = readTimeout
	}
}

func WithWriteTimeout(writeTimeout int16) func(conf *Neo4jConf) {
	return func(conf *Neo4jConf) {
		conf.WriteTimeout = writeTimeout
	}
}

func WithDbVersion(dbVersion string) func(conf *Neo4jConf) {
	return func(conf *Neo4jConf) {
		conf.DbVersion = dbVersion
	}
}

func WithRealm(realm string) func(conf *Neo4jConf) {
	return func(conf *Neo4jConf) {
		conf.Realm = realm
	}
}

type neo4jDbConn struct {
	database     string
	readTimeout  time.Duration
	writeTimeout time.Duration
	driver       neo4j.Driver
}

func (n *neo4jDbConn) getDefaultSession(accessMode neo4j.AccessMode, db string) (neo4j.Session, error) {
	sessionConfig := neo4j.SessionConfig{
		AccessMode: accessMode,
	}
	if len(db) > 0 {
		sessionConfig.DatabaseName = db
	}
	return n.driver.NewSession(sessionConfig), nil
}

func (n *neo4jDbConn) Exec(ctx context.Context, cypherInput *params_container.CypherInput, keys []string, accessMode neo4j.AccessMode) (*params_container.Result, error) {

	session, errGet := n.getDefaultSession(accessMode, n.database)
	if errGet != nil {
		return nil, errGet
	}
	defer session.Close()
	var re = &params_container.Result{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randI := r.Int63n(15) + 5
	errForRetry := utils.RetryWithDelay(5, time.Duration(randI)*time.Millisecond, func() error {
		result, err := session.Run(cypherInput.Cypher, cypherInput.Params, neo4j.WithTxTimeout(n.writeTimeout))
		if err != nil {
			return err
		}
		if nil != result.Err() {
			return result.Err()
		}
		if keys == nil {
			return nil
		}
		var reMap = make(map[string][]interface{})
		for result.Next() {
			record := result.Record()
			for index := range keys {
				key := keys[index]
				if value, ok := record.Get(key); ok {
					reMap[key] = append(reMap[key], value)
				}
			}
		}
		re.Data = reMap
		return nil
	})
	if nil != errForRetry {
		return nil, errForRetry
	}
	return re, nil
}

func (n *neo4jDbConn) TxExec(ctx context.Context, cypherInput *params_container.CypherInput, keys []string, accessMode neo4j.AccessMode) (*params_container.Result, error) {

	session, errGet := n.getDefaultSession(accessMode, n.database)
	if errGet != nil {
		return nil, errGet
	}
	defer session.Close()
	var re = &params_container.Result{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randI := r.Int63n(15) + 5
	errForRetry := utils.RetryWithDelay(5, time.Duration(randI)*time.Millisecond, func() error {
		_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			result, err := tx.Run(cypherInput.Cypher, cypherInput.Params)
			if err != nil {
				return nil, err
			}
			if nil != result.Err() {
				return nil, result.Err()
			}
			if keys == nil {
				return nil, nil
			}
			var reMap = make(map[string][]interface{})
			for result.Next() {
				record := result.Record()
				for index := range keys {
					key := keys[index]
					if value, ok := record.Get(key); ok {
						reMap[key] = append(reMap[key], value)
					}
				}
			}
			re.Data = reMap
			return re, nil
		}, neo4j.WithTxTimeout(n.writeTimeout))
		return err
	})
	if errForRetry != nil {
		return nil, errForRetry
	}
	return re, nil
}
