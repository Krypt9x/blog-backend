package repository

import (
	"context"
	"database/sql"

	"github.com/Krypt9x/blog-backend/internal/model/domain"
	"github.com/Krypt9x/blog-backend/internal/model/web"
)

type CommentRepository interface {
	Create(ctx context.Context, tx *sql.Tx, data domain.Comments) error
	GetById(ctx context.Context, tx *sql.Tx, id string) ([]web.CommentResponse, error)
	UpdateById(ctx context.Context, tx *sql.Tx, data domain.Comments) error
	Delete(ctx context.Context, tx *sql.Tx, id string) error
}
