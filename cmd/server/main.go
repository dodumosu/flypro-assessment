package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"flypro-assessment/internal/config"
	"flypro-assessment/internal/handlers"

	_ "github.com/joho/godotenv/autoload"
)

type APIServer struct {
	configuration config.ServerConfig
	server        *http.Server
	logger        *slog.Logger
}

func NewAPIServer(cfg config.ServerConfig, parentLogger *slog.Logger) (*APIServer, error) {
	return &APIServer{
		configuration: cfg,
		logger:        parentLogger.With("source", "server"),
		server:        &http.Server{},
	}, nil
}

func (s *APIServer) cleanup() {
	// cleanup hook
	s.logger.Info("cleaning up resources")
}

func (s *APIServer) Start() {
	address := net.JoinHostPort(s.configuration.Host, strconv.Itoa(s.configuration.Port))

	s.server.Addr = address

	serverErrors := make(chan error, 1)

	go func() {
		s.logger.Info("Starting HTTP server", "address", address)
		serverErrors <- s.server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		s.logger.Error("Server error", "error", err)
		os.Exit(1)
	case sig := <-shutdown:
		s.logger.Info("Received signal, starting graceful shutdown", "signal", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.server.Shutdown(ctx); err != nil {
			s.logger.Error("Server shutdown failed", "error", err)
		} else {
			s.logger.Info("Server gracefully stopped")
		}

		s.cleanup()
	}
}

func (s *APIServer) SetupHandler(handler http.Handler) {
	s.server.Handler = handler
}

func main() {
	settings := config.GetSettings()
	rootLogger := config.GetRootLogger(settings.Logging)

	apiServer, err := NewAPIServer(settings.Server, rootLogger)
	routerHandler := handlers.NewRouteHandler(rootLogger)

	// TODO: would it be better to inject it to server creation?
	router := routerHandler.SetupRoutes()
	apiServer.SetupHandler(router)

	if err != nil {
		panic(err)
	}

	apiServer.Start()
}
