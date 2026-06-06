// HTTP server and client implementation with context handling and graceful shutdown
// This code demonstrates how to create an HTTP server that can handle requests with context
// cancellation, and how to create an HTTP client that uses context for request management.
// It also includes a graceful shutdown mechanism for the HTTP server to ensure that it can cleanly
// shut down when it receives an interrupt signal, allowing ongoing requests to complete before the server stops.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Main function to start the HTTP server
func httpServerFund() {
	// Create a new ServeMux
	mux := http.NewServeMux()
	// Register a simple handler for the root path
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, HTTP Server!")
	})
	// Register the slow handler for the root path
	mux.HandleFunc("/slow", slowHandler)
	// Create server with timeouts
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// Start the server
	server.ListenAndServe()
}

// Slow handler to demonstrate context cancellation
func slowHandler(w http.ResponseWriter, r *http.Request) {
	// Get the context from the request
	ctx := r.Context()
	// Simulate a long-running process
	select {
	// Simulate work for 5 seconds
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "Finished processing")
	// Handle context cancellation
	case <-ctx.Done():
		fmt.Fprintf(w, "Request cancelled")
	}
}

// HTTP client function to demonstrate context usage
func httpClientFund(ctx context.Context, url string) (string, error) {
	// Create a new HTTP request with the provided context
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}
	// Create an HTTP client and perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	// Ensure the response body is closed after reading
	defer resp.Body.Close()
	return fmt.Sprintf("Response status: %s", resp.Status), nil
}

// Graceful shutdown of the HTTP server
func httpServerWithGracefulShutdown() {
	// Create a new ServeMux and register handlers
	mux := http.NewServeMux()
	// Register a simple handler for the root path
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, HTTP Server with Graceful Shutdown!")
	})
	// Create the HTTP server
	server := &http.Server{Addr: ":8080", Handler: mux}

	// Start the server in a separate goroutine
	go func() {
		// ListenAndServe returns an error when the server is shut down, so we check for that
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// Log the error if it's not due to the server being closed
			log.Fatal(err)
		}
	}()

	// Wait for an interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	// Notify the quit channel on SIGINT and SIGTERM signals
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// Block until a signal is received
	<-quit

	fmt.Println("Server shutting down...")

	// Create a context with a timeout for the shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Ensure the cancel function is called to release resources
	defer cancel()

	// Attempt to gracefully shut down the server with the context
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Error during server shutdown:", err)
	}

	fmt.Println("Server stopped gracefully")
}
