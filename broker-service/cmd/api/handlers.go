package main

import (
	"log"
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	log.Printf("CREATED THE PAYLOAD")
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the Broker",
	}
	_ = app.writeJSON(w, http.StatusOK, payload)

}
