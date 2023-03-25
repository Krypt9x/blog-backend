package repository

import (
	"context"
	"database/sql"

	"github.com/Krypt9x/blog-backend/internal/model/domain"
)

type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user domain.User) error
	Login(ctx context.Context, tx *sql.Tx, user domain.User) error
}
