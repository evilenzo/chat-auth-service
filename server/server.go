package server

import (
	dbo "main/db_operator"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

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
