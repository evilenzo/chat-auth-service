package server

import (
	"encoding/json"
	"log"
	"net/http"
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

func (s *Server) NameExists(w http.ResponseWriter, r *http.Request) {
	reqBody := struct {
		Name string `json:"name"`
	}{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqBody)
	hasErr := checkAndRespond(&w, "NameExists", "json parse error", err)
	if hasErr {
		return
	}

	rawResponse := struct {
		Exists bool `json:"exists"`
	}{}
	rawResponse.Exists, err = s.db.NameExists(reqBody.Name)

	serResponse, err := json.Marshal(rawResponse)
	hasErr = checkAndRespond(&w, "GET NameExists", "json serialize error", err)

	bytes, err := w.Write(serResponse)
	log.Printf("Respond %v bytes for `NameExists`", bytes)
	check("GET NameExists", "response writing", err)
}
