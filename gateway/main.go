package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Define the server configuration
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler),
	}

	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Print the request method (GET, POST, etc.)
	fmt.Println("Method:", r.Method)

	// Print the request URL
	fmt.Println("URL:", r.URL)

	// Print the request headers
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}

	// Read the request body and print its content
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
	}
	defer r.Body.Close() // Close the request body when we're done
	fmt.Println("Body:", string(body))
}
