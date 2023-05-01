package server

import (
	"log"
	"os"

	"github.com/eduardotvn/projeto-api/src/router"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func StartServer() Server {
	return Server{
		port:   os.Getenv("PORT"),
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := router.CreateRouter(s.server)

	log.Fatal(router.Run(":" + s.port))
}
