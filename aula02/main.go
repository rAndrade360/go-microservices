package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rAndrade360/go-microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "my-api", log.LstdFlags)
	hh := handlers.NewHelloHandler(l)

	sm := http.NewServeMux()

	sm.Handle("/", hh)

	http.ListenAndServe(":3000", sm)
}
