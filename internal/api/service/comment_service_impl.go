package services

import (
	"context"
	"database/sql"

	"github.com/Krypt9x/blog-backend/internal/api/repository"
	"github.com/Krypt9x/blog-backend/internal/model/domain"
	"github.com/Krypt9x/blog-backend/internal/model/web"
	"github.com/Krypt9x/blog-backend/pkg/helper"
)

type CommentServiceImpl struct {
	DB         *sql.DB
	Repository repository.CommentRepository
}

func NewCommentService(db *sql.DB, repository repository.CommentRepository) CommentService {
	return &CommentServiceImpl{
		DB:         db,
		Repository: repository,
	}
}

func (service *CommentServiceImpl) Create(ctx context.Context, data domain.Comments) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	err = service.Repository.Create(ctx, tx, data)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx, err)
	return nil
}

func (service *CommentServiceImpl) GetById(ctx context.Context, id string) []web.CommentResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	allData, err := service.Repository.GetById(ctx, tx, id)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx, err)
	return allData
}

func (service *CommentServiceImpl) UpdateById(ctx context.Context, data domain.Comments) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	err = service.Repository.UpdateById(ctx, tx, data)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx, err)
	return nil
}

func (service *CommentServiceImpl) Delete(ctx context.Context, id string) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	err = service.Repository.Delete(ctx, tx, id)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx, err)
	return nil
}
