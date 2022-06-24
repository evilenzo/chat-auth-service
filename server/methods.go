package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func checkAndRespond(w *http.ResponseWriter, method string, sub string, err error) bool {
	if err != nil {
		log.Printf("Error in %v (%v): %v", method, sub, err)
		http.Error(*w, sub, http.StatusInternalServerError)

		return true
	}

	return false
}

// REST API

func (s *Server) NameExists(c *gin.Context) {
	body := struct {
		Name string `json:"name"`
	}{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	if exists, err := s.db.NameExists(body.Name); err != nil {
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
