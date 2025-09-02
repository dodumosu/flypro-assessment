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
	"flypro-assessment/internal/services"
	"flypro-assessment/internal/utils"

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

func (s *APIServer) Start(shutdown <-chan os.Signal, done chan<- struct{}) {
	address := net.JoinHostPort(s.configuration.Host, strconv.Itoa(s.configuration.Port))

	s.server.Addr = address

	serverErrors := make(chan error, 1)

	go func() {
		s.logger.Info("Starting HTTP server", "address", address)
		serverErrors <- s.server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		s.logger.Error("Server error", "error", err)
		os.Exit(1)
	case sig := <-shutdown:
		close(done) // Signal ticker to stop
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

	// create rates service
	cache, err := services.NewRedisCache(settings.Redis)
	if err != nil {
		panic(err)
	}
	defer cache.Close()
	ratesClient := utils.NewCurrencyRatesClient(settings.RatesConfig)
	ratesService := services.NewCurrencyRatesService(cache, ratesClient)

	// Load rates on startup
	ctx := context.Background()
	if err := ratesService.LoadRates(ctx); err != nil {
		rootLogger.Error("Failed to load initial currency rates", "error", err)
	} else {
		rootLogger.Info("Initial currency rates loaded successfully")
	}

	apiServer, err := NewAPIServer(settings.Server, rootLogger)
	if err != nil {
		panic(err)
	}

	routerHandler := handlers.NewRouteHandler(rootLogger)

	// TODO: would it be better to inject it to server creation?
	router := routerHandler.SetupRoutes()
	apiServer.SetupHandler(router)

	// Set up rates loading ticker (every 6 hours)
	ticker := time.NewTicker(6 * time.Hour)

	shutdown := make(chan os.Signal, 1)
	done := make(chan struct{})
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Start ticker goroutine
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				ctx := context.Background()
				if err := ratesService.LoadRates(ctx); err != nil {
					rootLogger.Error("Failed to load currency rates", "error", err)
				} else {
					rootLogger.Info("Currency rates refreshed successfully")
				}
			case <-done:
				rootLogger.Info("Stopping rates ticker")
				return
			}
		}
	}()

	apiServer.Start(shutdown, done)
}
