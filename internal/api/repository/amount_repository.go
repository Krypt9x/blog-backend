package repository

import (
	"context"
	"database/sql"
)

type AmountRepository interface {
	Create(ctx context.Context, tx *sql.Tx, id string) error
	GetAmountCommentsById(ctx context.Context, tx *sql.Tx, id string) (uint64, error)
	GetAmountViewsById(ctx context.Context, tx *sql.Tx, id string) (uint64, error)
	UpdateAmountCommentsById(ctx context.Context, tx *sql.Tx, id string) (uint64, error)
	UpdateAmountViewsById(ctx context.Context, tx *sql.Tx, id string) (uint64, error)
	Delete(ctx context.Context, tx *sql.Tx, id string) error
}
