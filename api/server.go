package api

import (
	db "github.com/Manan-Rastogi/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

//Serves HTTP request for our banking service.
type Server struct {
	store *db.Store
	router *gin.Engine
}

// Creates a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("accounts/:id", server.getAccount)
	router.GET("accounts", server.listAccount)

	server.router = router
	return server
}

// Runs server on specified address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Return the error in map[string]interface{} format.
func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(), 
	}	
}
