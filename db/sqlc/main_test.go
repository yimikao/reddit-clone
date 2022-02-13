package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/yimikao/reddit-clone/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	c, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("couldn't load config: ", err)
	}

	testDB, err = sql.Open(c.DBDriver, c.DBSource)
	if err != nil {
		log.Fatal("couldn't connect to database: ", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
