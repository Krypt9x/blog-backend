package services

import (
	"context"

	"github.com/Krypt9x/blog-backend/internal/model/domain"
)

type UserService interface {
	Register(ctx context.Context, user domain.User) error
	Login(ctx context.Context, user domain.User) error
}
