package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rAndrade360/go-microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "my-api", log.LstdFlags)
	hh := handlers.NewHelloHandler(l)
	gh := handlers.NewGoodByeHandler(l)

	sm := http.NewServeMux()

	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Handler:      sm,
		Addr:         ":3000",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		ErrorLog:     l,
	}

	s.ListenAndServe()
}
