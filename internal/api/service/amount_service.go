package services

import "context"

type AmountService interface {
	Create(ctx context.Context, id string) error
	GetAmountCommentById(ctx context.Context, id string) (uint64, error)
	GetAmountViewsById(ctx context.Context, id string) (uint64, error)
	UpdateAmountCommentsById(ctx context.Context, id string) (uint64, error)
	UpdateAmountViewsById(ctx context.Context, id string) (uint64, error)
	Delete(ctx context.Context, id string) error
}
