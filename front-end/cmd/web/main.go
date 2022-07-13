package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	// The documentation for ServeMux explains how patterns are matched.
	// ServeMux is an HTTP request multiplexer.
	// ServeMux matches the URL of each incoming request against a list of registered patterns and calls the
	// handler for the pattern that most closely matches the URL.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
		render(w, "test.page.gohtml")
	})

	fmt.Println("Starting front end service on port 80. Lets log a couple of things :)")
	// ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
	// Accepted connections are configured to enable TCP keep-alives.
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {

	partials := []string{
		"./cmd/web/templates/base.layout.gohtml",
		"./cmd/web/templates/header.partial.gohtml",
		"./cmd/web/templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	// ParseFiles creates a new Template and parses the template definitions from the named files.
	// The returned template's name will have the (base) name and (parsed) contents of the first file.
	// There must be at least one file. If an error occurs, parsing stops and the returned *Template is nil.
	// Return Value: (*Template, error) Template is a specialized Template from "text/template" that produces a safe HTML document fragment.
	// Template is a type with its own methods.
	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute applies a parsed template to the specified data object, writing the output to wr.
	// If an error occurs executing the template or writing its output, execution stops,
	//but partial results may already have been written to the output writer.
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
