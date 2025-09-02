package handlers

import (
	"context"

	"flypro-assessment/internal/dto"

	"github.com/danielgtaylor/huma/v2"
)

func (r *RouteHandler) CreateExpense(ctx context.Context, input *dto.CreateExpenseReportRequest) (*dto.ExpenseCreateResponse, error) {
	return nil, huma.Error501NotImplemented("not yet implemented")
}

func (r *RouteHandler) ListExpenses(ctx context.Context, input *dto.ExpenseListRequest) (*dto.ExpenseListResponse, error) {
	return nil, huma.Error501NotImplemented("not yet implemented")
}

func (r *RouteHandler) GetExpense(ctx context.Context, input *dto.ExpenseDetailRequest) (*dto.ExpenseDetailResponse, error) {
	return nil, huma.Error501NotImplemented("not yet implemented")
}

func (r *RouteHandler) UpdateExpense(ctx context.Context, input *dto.ExpenseUpdateRequest) (*dto.ExpenseDetailResponse, error) {
	return nil, huma.Error501NotImplemented("not yet implemented")
}

func (r *RouteHandler) DeleteExpense(ctx context.Context, input *dto.ExpenseDetailRequest) (*struct{}, error) {
	return nil, huma.Error501NotImplemented("not yet implemented")
}
