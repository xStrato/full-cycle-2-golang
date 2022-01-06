package webapi

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer(p string) *Server {
	return &Server{
		port:   p,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	log.Println("Server is running on port: " + s.port)
	log.Fatalln(s.server.Run(":" + s.port))
}
func (cr *Server) Router(f func(*gin.Engine) error) error {
	if err := f(cr.server); err != nil {
		return err
	}
	return nil
}
