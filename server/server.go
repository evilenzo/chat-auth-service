package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	dbo "main/db_operator"

	"github.com/gorilla/mux"
	"net/http"
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
	gracefulWait := time.Second * 15

	router := mux.NewRouter()
	router.StrictSlash(false)

	router.HandleFunc("/name_exists", s.NameExists).Methods("GET")

	srv := &http.Server{
		Addr: ":8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		//WriteTimeout: time.Second * 15,
		//ReadTimeout:  time.Second * 15,
		//IdleTimeout:  time.Second * 60,
		Handler: router,
	}

	go func() {
		log.Println("Server initialized")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), gracefulWait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err := srv.Shutdown(ctx)
	check("server run", "server shutdown", err)

	log.Println("shutting down")
	os.Exit(0)
}
