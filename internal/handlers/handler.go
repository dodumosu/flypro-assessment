package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"flypro-assessment/internal/config"
	"flypro-assessment/internal/dto"
	"flypro-assessment/internal/middleware"

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

	// set up middleware
	middleware.SetupMiddleware(router, r.logger, settings)

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

	userManagementGroup := huma.NewGroup(api)
	huma.Register(userManagementGroup, huma.Operation{
		Description: "Create user",
		Method:      http.MethodPost,
		OperationID: "create-user",
		Path:        UserCreatePath,
		Summary:     "Create user",
		Tags:        []string{"Users"},
	}, r.CreateUser)
	huma.Register(userManagementGroup, huma.Operation{
		Description: "Get user",
		Method:      http.MethodGet,
		OperationID: "get-user",
		Path:        UserCreatePath,
		Summary:     "Get user",
		Tags:        []string{"Users"},
	}, r.GetUser)

	expenseManagementGroup := huma.NewGroup(api)
	huma.Register(expenseManagementGroup, huma.Operation{
		Description: "Create expense",
		Method:      http.MethodPost,
		OperationID: "create-expense",
		Path:        ExpenseCreatePath,
		Summary:     "Create expense",
		Tags:        []string{"Expenses"},
	}, r.CreateExpense)
	huma.Register(expenseManagementGroup, huma.Operation{
		Description: "List expenses",
		Method:      http.MethodGet,
		OperationID: "list-expenses",
		Path:        ExpenseListPath,
		Summary:     "List expenses",
		Tags:        []string{"Expenses"},
	}, r.ListExpenses)
	huma.Register(expenseManagementGroup, huma.Operation{
		Description: "Get expense",
		Method:      http.MethodGet,
		OperationID: "get-expense",
		Path:        ExpenseDetailPath,
		Summary:     "Get expense",
		Tags:        []string{"Expenses"},
	}, r.GetExpense)
	huma.Register(expenseManagementGroup, huma.Operation{
		Description: "Update expense",
		Method:      http.MethodPut,
		OperationID: "update-expense",
		Path:        ExpenseUpdatePath,
		Summary:     "Update expense",
		Tags:        []string{"Expenses"},
	}, r.UpdateExpense)
	huma.Register(expenseManagementGroup, huma.Operation{
		Description: "Delete expense",
		Method:      http.MethodDelete,
		OperationID: "delete-expense",
		Path:        ExpenseDeletePath,
		Summary:     "Delete expense",
		Tags:        []string{"Expenses"},
	}, r.DeleteExpense)

	expenseReportGroup := huma.NewGroup(api)
	huma.Register(expenseReportGroup, huma.Operation{
		Description: "Create expense report",
		Method:      http.MethodPost,
		OperationID: "create-expense-report",
		Path:        ReportCreatePath,
		Summary:     "Create expense report",
		Tags:        []string{"Expense reports"},
	}, r.CreateExpenseReport)
	huma.Register(expenseReportGroup, huma.Operation{
		Description: "Add expenses to report",
		Method:      http.MethodPost,
		OperationID: "add-expenses-to-report",
		Path:        ReportExpenseAddPath,
		Summary:     "Add expenses to report",
		Tags:        []string{"Expense reports"},
	}, r.AddExpensesToReport)
	huma.Register(expenseReportGroup, huma.Operation{
		Description: "List reports",
		Method:      http.MethodGet,
		OperationID: "list-reports",
		Path:        ReportListPath,
		Summary:     "List reports",
		Tags:        []string{"Expense reports"},
	}, r.ListExpenseReports)
	huma.Register(expenseReportGroup, huma.Operation{
		Description: "Submit report for approval",
		Method:      http.MethodPut,
		OperationID: "submit-report-for-approval",
		Path:        ReportListPath,
		Summary:     "Submit report for approval",
		Tags:        []string{"Expense reports"},
	}, r.SubmitReportForApproval)

	return router
}
