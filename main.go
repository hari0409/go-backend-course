package main

import (
	"database/sql"
	"log"

	"github.com/hari0409/backend-go/api"
	db "github.com/hari0409/backend-go/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://postgres:mysecretpassword@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Println(err)
	}

	server := api.NewServer(db.NewStore(conn))
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
