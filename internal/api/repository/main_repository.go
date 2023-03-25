package repository

import (
	"context"
	"database/sql"

	"github.com/Krypt9x/blog-backend/internal/model/domain"
	"github.com/Krypt9x/blog-backend/internal/model/web"
)

type MainRepository interface {
	Create(ctx context.Context, tx *sql.Tx, data domain.MainDomain) (web.MainModelResponse, error)
	GetAll(ctx context.Context, tx *sql.Tx) []web.MainModelResponse
	GetByUsername(ctx context.Context, tx *sql.Tx, username string) []web.MainModelResponse
	GetById(ctx context.Context, tx *sql.Tx, id string) (web.MainModelResponse, error)
	UpdateById(ctx context.Context, tx *sql.Tx, data domain.MainDomain) (web.MainModelResponse, error)
	Delete(ctx context.Context, tx *sql.Tx, id string) error
}
