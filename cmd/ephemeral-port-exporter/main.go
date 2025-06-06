package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/davidsugianto/ephemeral-port-exporter/internal/collector"
	"github.com/davidsugianto/ephemeral-port-exporter/internal/router"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	prometheus.MustRegister(collector.NewEphemeralPortCollector())

	port := os.Getenv("PORT")
	if port == "" {
		port = "2112"
	}

	r := router.New()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Graceful shutdown
	go func() {
		log.Printf("Exporter starting on port %s\n", port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down exporter...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Exporter forced to shutdown: %v", err)
	}

	log.Println("Exporter exited")
}
