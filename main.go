package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jordbick/Practice-App/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("", hh)

	// Constructs a HTTP server and registers a default handler to it
	http.ListenAndServe(":9090", sm)

}
