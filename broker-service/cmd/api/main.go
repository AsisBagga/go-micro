package main

import (
	"fmt"
	"log"
	"net/http"
)

// We want to have one route. Accept JSON Payload and respond. Dead Simple.
// Over time this service will get more complex
// We will need a routeing service.

const webPort = "80" //Docker can listen on container any port

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	// Define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	log.Printf("HTTP SERVER IS DEFINED")
	// Start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
