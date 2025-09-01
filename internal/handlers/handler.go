package handlers

import (
	"context"
	"flypro-assessment/internal/config"
	"flypro-assessment/internal/dto"
	"log/slog"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	logger *slog.Logger
}

func NewRouteHandler(logger *slog.Logger) *RouteHandler {
	return &RouteHandler{
		logger: logger,
	}
}

func (r *RouteHandler) healthCheck(ctx context.Context, input *struct{}) (*dto.HealthCheckDTOEnvelope, error) {
	response := dto.HealthCheckDTO{
		Status: "ok",
	}

	return &dto.HealthCheckDTOEnvelope{
		Body: response,
	}, nil
}

func (r *RouteHandler) SetupRoutes() http.Handler {
	router := gin.New()
	settings := config.GetSettings()

	docsConfig := huma.DefaultConfig(settings.APISettings.Description, settings.APISettings.Version)
	docsConfig.DocsPath = DocsPath
	api := humagin.New(router, docsConfig)

	huma.Register(api, huma.Operation{
		Description: "Health check",
		Method:      http.MethodGet,
		OperationID: "health-check",
		Path:        HealthCheckPath,
		Summary:     "Health check",
		Tags:        []string{"System"},
	}, r.healthCheck)

	return router
}
