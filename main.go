package main

import (
	"database/sql"
	"log"

	"github.com/hari0409/backend-go/api"
	db "github.com/hari0409/backend-go/db/sqlc"
	"github.com/hari0409/backend-go/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error Opening DB Connection: ", err)
	}

	server := api.NewServer(db.NewStore(conn))
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
