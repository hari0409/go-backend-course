package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/hari0409/backend-go/db/sqlc"
)

type Server struct {
	store db.Store // To execute the DB queries and the transaction.
	// We dont use the *db.Queries here because Queries is already in the Store.
	router *gin.Engine // Router of gin to send the request to the handler.
}

func NewServer(store db.Store) *Server {
	// Create the server & the router
	server := &Server{
		store: store,
	}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	// Add the Routes to the router.
	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponce(err error) gin.H {
	return gin.H{"error": err.Error()}
}
