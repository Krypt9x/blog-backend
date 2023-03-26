package services

import (
	"context"
	"database/sql"

	"github.com/Krypt9x/blog-backend/internal/api/repository"
	"github.com/Krypt9x/blog-backend/internal/model/domain"
	"github.com/Krypt9x/blog-backend/pkg/helper"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	DB         sql.DB
	Repository repository.UserRepository
}

func NewUsersService(DB *sql.DB, repository repository.UserRepository) UserService {
	return &UserServiceImpl{
		DB:         *DB,
		Repository: repository,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, user domain.User) error {
	userData := domain.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	err := validator.New().Struct(userData)
	if err != nil {
		return err
	}

	dbTx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	err = service.Repository.Register(ctx, dbTx, userData)
	if err != nil {
		return err
	}

	defer helper.CommitOrRollback(dbTx, err)
	return nil
}

func (service *UserServiceImpl) Login(ctx context.Context, user domain.User) error {
	userData := domain.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	dbTx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	err = service.Repository.Login(ctx, dbTx, userData)
	if err != nil {
		return err
	}

	defer helper.CommitOrRollback(dbTx, err)
	return nil
}
