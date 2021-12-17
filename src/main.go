package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rAndrade360/go-microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "my-api", log.LstdFlags)
	ph := handlers.NewProductHandler(l)

	sm := http.NewServeMux()

	sm.Handle("/", ph)

	s := &http.Server{
		Handler:      sm,
		Addr:         ":3000",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		ErrorLog:     l,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Received terminated gracefull shutdown ", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
