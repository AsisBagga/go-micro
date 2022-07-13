package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// we will have 3 functions over here.
// one for reading json
// one for writing json
// one for error json

func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576 // one MB

	// MaxBytesReader's result is a ReadCloser, returns a non-EOF error for a Read beyond the limit,
	// and closes the underlying reader when its Close method is called.
	// MaxBytesReader prevents clients from accidentally or maliciously sending a large request and wasting server resources.
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// Returns Decoder
	dec := json.NewDecoder(r.Body)
	// A Decoder reads and decodes JSON values from an input stream.
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("Body must have only a single json value")
	}
	return nil
}

func (app *Config) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// writing headers
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	w.Header().Set("Content-Type", "application/json")
	// writing status
	w.WriteHeader(status)
	// writing data
	_, err = w.Write(out)

	return err
}

func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()
	return app.writeJSON(w, statusCode, payload)
}
