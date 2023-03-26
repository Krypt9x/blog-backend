package services

import (
	"github.com/Krypt9x/blog-backend/internal/api/repository"
	"github.com/Krypt9x/blog-backend/pkg/helper"

	"context"
	"database/sql"
)

type AmountServiceImpl struct {
	DB         *sql.DB
	Repository repository.AmountRepository
}

func NewAmountService(DB *sql.DB, amountRepository repository.AmountRepository) AmountService {
	return &AmountServiceImpl{
		DB:         DB,
		Repository: amountRepository,
	}
}

func (service *AmountServiceImpl) Create(ctx context.Context, id string) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	err = service.Repository.Create(ctx, tx, id)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx, err)
	return nil
}

func (service *AmountServiceImpl) GetAmountCommentById(ctx context.Context, id string) (uint64, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	data, err := service.Repository.GetAmountCommentsById(ctx, tx, id)
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(tx, err)
	return data, nil
}

func (service *AmountServiceImpl) GetAmountViewsById(ctx context.Context, id string) (uint64, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	data, err := service.Repository.GetAmountViewsById(ctx, tx, id)
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(tx, err)
	return data, nil
}

func (service *AmountServiceImpl) UpdateAmountCommentsById(ctx context.Context, id string) (uint64, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	updateAmountComments, err := service.Repository.UpdateAmountCommentsById(ctx, tx, id)
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(tx, err)
	return updateAmountComments, nil
}

func (service *AmountServiceImpl) UpdateAmountViewsById(ctx context.Context, id string) (uint64, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	updatedAmountViews, err := service.Repository.UpdateAmountViewsById(ctx, tx, id)
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(tx, err)
	return updatedAmountViews, nil
}

func (service *AmountServiceImpl) Delete(ctx context.Context, id string) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	err = service.Repository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx, err)
	return nil
}
