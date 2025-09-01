package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"flypro-assessment/internal/config"

	_ "github.com/joho/godotenv/autoload"
)

type APIServer struct {
	configuration config.ServerConfig
	server        *http.Server
}

func NewAPIServer(cfg config.ServerConfig) (*APIServer, error) {
	return &APIServer{
		configuration: cfg,
		server:        &http.Server{},
	}, nil
}

func (s *APIServer) cleanup() {
	// cleanup hook
	fmt.Printf("cleaning up resources\n")
}

func (s *APIServer) Start() {
	address := net.JoinHostPort(s.configuration.Host, strconv.Itoa(s.configuration.Port))

	s.server.Addr = address

	serverErrors := make(chan error, 1)

	go func() {
		fmt.Printf("Starting HTTP server on: %v\n", address)
		serverErrors <- s.server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		fmt.Printf("Server error: %v\n", err)
		os.Exit(1)
	case sig := <-shutdown:
		fmt.Printf("Received signal: %v, message: %s\n", sig, "Starting graceful shutdown...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.server.Shutdown(ctx); err != nil {
			fmt.Printf("Server shutdown failed: %v\n", err)
		} else {
			fmt.Printf("Server gracefully stopped\n")
		}

		s.cleanup()
	}
}

func main() {
	settings := config.GetSettings()
	apiServer, err := NewAPIServer(settings.Server)
	if err != nil {
		panic(err)
	}

	apiServer.Start()
}
