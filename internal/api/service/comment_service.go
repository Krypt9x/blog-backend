package services

import (
	"context"

	"github.com/Krypt9x/blog-backend/internal/model/domain"
	"github.com/Krypt9x/blog-backend/internal/model/web"
)

type CommentService interface {
	Create(ctx context.Context, data domain.Comments) error
	GetById(ctx context.Context, id string) []web.CommentResponse
	UpdateById(ctx context.Context, data domain.Comments) error
	Delete(ctx context.Context, id string) error
}
