package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// specify who is allowed to connect
	// setting up middleware
	log.Printf("In the router right now")
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true, //allowed to deal with credentials requests
		MaxAge:           300,
	}))

	// To allow to check this service for a heartbeat
	mux.Use(middleware.Heartbeat("/ping"))
	log.Printf("mux is routing the request to app.Broker")
	mux.Post("/", app.Broker)
	return mux
}

//Okay, so this will be a route, but I want to actually add a receiver here that allows me to share
//any configuration I might have from my application with routes in case I need them.
