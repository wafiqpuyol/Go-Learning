package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wafiqpuyol/Go-Learning/api/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()
	fmt.Println(*cfg)

	// start db

	// start router
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})
	server := http.Server{
		Addr:    cfg.HTTP_Server.Address,
		Handler: router,
	}

	// start server & graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	select {
	case <-done:
		slog.Info("shutting down server", slog.String("address", cfg.HTTP_Server.Address))
		break
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shut down", slog.String("error", err.Error()))
	}
}
