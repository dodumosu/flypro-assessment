package handlers

import (
	"context"

	"flypro-assessment/internal/dto"

	"github.com/danielgtaylor/huma/v2"
)

func (r *RouteHandler) CreateUser(ctx context.Context, input *dto.UserCreateRequest) (*dto.UserCreateResponse, error) {
	return nil, huma.Error501NotImplemented("not yet implemented")
}

func (r *RouteHandler) GetUser(ctx context.Context, input *dto.UserDetailRequest) (*dto.UserDetailResponse, error) {
	return nil, huma.Error501NotImplemented("not yet implemented")
}
