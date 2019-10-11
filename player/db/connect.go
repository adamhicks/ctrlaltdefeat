package db

import (
	"database/sql"
	"flag"
	"runtime"
	"strings"
	"testing"

	"github.com/corverroos/unsure"
	"github.com/luno/jettison/log"
)

var (
	dbURI = flag.String("player_db", "mysql://root@unix("+unsure.SockFile()+")/player?",
		"player DB URI")
)

type PlayerDB struct {
	DB        *sql.DB
	ReplicaDB *sql.DB
}

// ReplicaOrMaster returns the replica DB if available, otherwise the master.
func (db *PlayerDB) ReplicaOrMaster() *sql.DB {
	if db.ReplicaDB != nil {
		return db.ReplicaDB
	}
	return db.DB
}

func Connect(p string) (*PlayerDB, error) {
	appendPlayerToURI(p)

	ok, err := unsure.MaybeRecreateSchema(*dbURI, getSchemaPath())
	if err != nil {
		return nil, err
	} else if ok {
		log.Info(nil, "recreated schema")
	}

	dbc, err := unsure.Connect(*dbURI)
	if err != nil {
		return nil, err
	}
	return &PlayerDB{
		DB:        dbc,
		ReplicaDB: dbc,
	}, nil
}

func appendPlayerToURI(p string) {
	pdb := flag.Lookup("player_db")
	uri := pdb.Value.String() + p
	_ = flag.Set("player_db", uri)
}

func ConnectForTesting(t *testing.T) *sql.DB {
	return unsure.ConnectForTesting(t, getSchemaPath())
}

func getSchemaPath() string {
	_, filename, _, _ := runtime.Caller(0)
	return strings.Replace(filename, "connect.go", "schema.sql", 1)
}
