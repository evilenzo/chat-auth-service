package server

import (
	dbo "main/db_operator"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func check(method string, sub string, err error) {
	if err != nil {
		log.Printf("Error in %v (%v): %v", method, sub, err)
	}
}

type Server struct {
	db dbo.DatabaseOperator
}

func CreateServer(operator dbo.DatabaseOperator) Server {
	return Server{operator}
}

func (s *Server) Run() {
	r := gin.Default()

	r.GET("/name_exists", s.NameExists)

	err := r.Run()
	if err != nil {
		log.Panic("Start server error: ", err)
	}
}
