package server

import (
	"errors"
	dbe "main/db_operator/db_err"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// REST API

func (s *Server) NameExists(c *gin.Context) {
	body := struct {
		Name string `json:"name"`
	}{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if exists, err := s.db.NameExists(c, body.Name); err != nil {
		c.Status(http.StatusInternalServerError)
		log.Error("Error during /name_exists response: ", err)
	} else {
		response := struct {
			Exists bool `json:"exists"`
		}{
			Exists: exists,
		}

		c.JSON(http.StatusOK, response)
		log.Trace("Successful /name_exists response")
	}
}

func (s *Server) AuthApp(c *gin.Context) {
	body := struct {
		ID       int64  `json:"id"`
		Password string `json:"password"`
	}{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if token, err := s.db.AuthApp(c, body.ID, body.Password); err != nil {
		if errors.Is(err, dbe.WrongPassword) {
			c.String(http.StatusForbidden, "Wrong password")
		} else if errors.Is(err, dbe.RecordNotFound) {
			c.String(http.StatusForbidden, "User with this id not found")
		} else {
			c.Status(http.StatusInternalServerError)
			log.Error("Error during /auth_app response: ", err)
		}
	} else {
		response := struct {
			Token string `json:"token"`
		}{token}
		c.JSON(http.StatusOK, &response)
		log.Trace("Successful /auth_app response")
	}
}
