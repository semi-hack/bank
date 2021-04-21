package routes

import (
	"github.com/gin-gonic/gin"
	db "bank/db/sqlc"
)

// Server ...
type Server struct {
	store *db.Store
	router *gin.Engine
}

// Initialize a new http server 
func Initialize(store *db.Store) *Server {
	server := &Server{store: store}
	r := gin.Default()

	server.router = r
	return server

	//r.POST("https://moneywave.herokuapp.com/v1/merchant/verify/", controllers)
}