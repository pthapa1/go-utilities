package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Port we listen on.
const portNum string = ":8080"

// Info handles the /callback endpoint, prints the captured code and shuts down the server.
func Info(w http.ResponseWriter, r *http.Request) {
	// Extract the 'code' query parameter from the URL
	code := r.URL.Query().Get("code")
	if code != "" {
		// Print the captured value
		fmt.Printf("Received code: %s\n", code)
		// Send response to the client
		fmt.Fprintf(w, "Code received: %s\n", code)

		// Shutdown the server gracefully
		log.Println("Shutting down the server.")
		os.Exit(0) // Exit the application
	} else {
		// If no code is provided, inform the user
		fmt.Fprintf(w, "No 'code' query parameter found.")
	}
}

// Server starts the HTTP server and listens for requests.
func Server() {
	log.Println("Starting our simple http server.")

	// Register the handler function for the /callback route
	http.HandleFunc("/callback", Info)

	log.Println("Started on port", portNum)
	fmt.Println("To close connection CTRL+C :-)")

	// Spinning up the server.
	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
