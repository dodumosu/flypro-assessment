package handlers

import (
	"context"

	"flypro-assessment/internal/dto"

	"github.com/danielgtaylor/huma/v2"
)

func (r *RouteHandler) CreateExpenseReport(ctx context.Context, input *dto.CreateExpenseReportRequest) (*dto.ExpenseReportDetailResponse, error) {
	return nil, huma.Error501NotImplemented("not yet implemented")
}

func (r *RouteHandler) AddExpensesToReport(ctx context.Context, input *dto.AddExpensesToReportRequest) (*dto.ExpenseReportDetailResponse, error) {
	return nil, huma.Error501NotImplemented("not yet implemented")
}

func (r *RouteHandler) ListExpenseReports(ctx context.Context, input *dto.ExpenseReportListRequest) (*dto.ExpenseReportListResponse, error) {
	return nil, huma.Error501NotImplemented("not yet implemented")
}

func (r *RouteHandler) SubmitReportForApproval(ctx context.Context, input *dto.ExpenseReportApprovalRequest) (*dto.ExpenseReportDetailResponse, error) {
	return nil, huma.Error501NotImplemented("not yet implemented")
}
