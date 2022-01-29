package main

import (
	"database/sql"
	"log"

	"github.com/yimikao/reddit-clone/api"
	db "github.com/yimikao/reddit-clone/db/sqlc"
	"github.com/yimikao/reddit-clone/util"
)

func main() {
	c, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("couldn't load config %v", err)
	}

	conn, err := sql.Open(c.DBDriver, c.DBSource)
	if err != nil {
		log.Fatalf("couldn't open database connection: %v", err)
	}

	s := db.NewStore(conn)
	svr, err := api.NewServer(c, s)
	if err != nil {
		log.Fatalf("couldn't create server: %v", err)
	}

	err = svr.Start(c.ServerAddr)
	if err != nil {
		log.Fatalf("couln't start server: %v", err)
	}
}
