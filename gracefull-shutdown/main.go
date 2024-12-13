package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Create a context to handle graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	// Create a WaitGroup to keep track of running goroutines
	var wg sync.WaitGroup

	// Start the HTTP server
	go startHTTPServer(ctx, &wg)

	// Listen for termination signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Wait for termination signal
	<-signalCh

	// Start the graceful shutdown process
	slog.Info("Gracefully shutting down HTTP server")

	// Cancel the context to signal the HTTP server to stop
	cancel()

	// Wait for the HTTP server to finish
	wg.Wait()

	slog.Info("Shutdown complete.")
}

func startHTTPServer(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.NewServeMux(),
	}

	// Start the HTTP server in a separate goroutine
	go func() {
		fmt.Println("Starting HTTP server...")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP server error: %s\n", err)
		}
	}()

	// Wait for the context to be canceled
	select {
	case <-ctx.Done():
		// Shutdown the server gracefully
		slog.Info("Shutting down HTTP server gracefully...")
		shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelShutdown()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			slog.Info("HTTP server shutdown error: ", slog.String("error", err.Error()))
		}
	}

	slog.Info("HTTP server stopped.")
}
