package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Port we listen on.
const portNum string = ":8080"

// Callback handles the /callback endpoint, prints the captured code, lets the user know they can close the window, and shuts down the server.
func Callback(w http.ResponseWriter, r *http.Request) {
	// Extract the 'code' query parameter from the URL
	code := r.URL.Query().Get("code")
	if code != "" {
		// Print the captured value in the terminal
		fmt.Printf("Received code: %s\n", code)

		// Respond to the client with the code and shutdown message
		fmt.Fprintf(w, "Code received: %s\n", code)
		fmt.Fprintf(w, "You can now safely close this window.\n")

		// Log the server shutdown
		log.Println("Shutting down the server.")

		// Gracefully shut down the server after sending the response
		go func() {
			// Wait a little to ensure the message reaches the browser before shutdown
			// (a small delay is often useful in some scenarios, but it's not strictly necessary)
			// time.Sleep(time.Second)
			os.Exit(0) // Exit the application
		}()
	} else {
		// If no code is provided, inform the user
		fmt.Fprintf(w, "No 'code' query parameter found.")
	}
}

// Server starts the HTTP server and listens for requests.
func Server() {
	log.Println("Starting our simple http server.")

	// Register the handler function for the /callback route
	http.HandleFunc("/callback", Callback)

	log.Println("Started on port", portNum)
	fmt.Println("To close connection CTRL+C :-)")

	// Spinning up the server.
	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
