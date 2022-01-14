package main

import (
	"fmt"
	"net/http"
	"nst-golang-course/04-html-server/handlers"
)

func main() {
	// Create a server request handler
	serveMux := http.NewServeMux()

	// set routes and route handlers
	serveMux.HandleFunc("/", handlers.RootHandler)

	// Create listener and add handler
	// err := http.ListenAndServe("localhost:8080", serverMux)
	if err := http.ListenAndServe("localhost:8080", serveMux); err != nil {
		fmt.Println("!ERROR: \n\t", err)
	}
}
