package services

import (
	"context"

	"github.com/Krypt9x/blog-backend/internal/model/domain"
	"github.com/Krypt9x/blog-backend/internal/model/web"
)

type MainService interface {
	Create(ctx context.Context, data domain.MainDomain) web.MainModelResponse
	GetAll(ctx context.Context) []web.MainModelResponse
	GetByUsername(ctx context.Context, username string) []web.MainModelResponse
	GetById(ctx context.Context, id string) (web.MainModelResponse, error)
	UpdateById(ctx context.Context, data domain.MainDomain) (web.MainModelResponse, error)
	Delete(ctx context.Context, id string) error
}
