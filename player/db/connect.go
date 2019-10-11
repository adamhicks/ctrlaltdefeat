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
	dbURI = flag.String("player_db", "mysql://root@unix("+unsure.SockFile()+")/player_",
		"player DB URI")
)

func Connect(p string) (*sql.DB, error) {
	uri := *dbURI + p + "?"

	ok, err := unsure.MaybeRecreateSchema(uri, getSchemaPath())
	if err != nil {
		return nil, err
	} else if ok {
		log.Info(nil, "recreated schema")
	}

	return unsure.Connect(uri)
}

func ConnectForTesting(t *testing.T) *sql.DB {
	return unsure.ConnectForTesting(t, getSchemaPath())
}

func getSchemaPath() string {
	_, filename, _, _ := runtime.Caller(0)
	return strings.Replace(filename, "connect.go", "schema.sql", 1)
}
