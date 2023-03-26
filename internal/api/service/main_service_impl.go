package services

import (
	"context"
	"database/sql"
	"log"

	"github.com/Krypt9x/blog-backend/internal/api/repository"
	"github.com/Krypt9x/blog-backend/internal/model/domain"
	"github.com/Krypt9x/blog-backend/internal/model/web"
	"github.com/Krypt9x/blog-backend/pkg/helper"
)

type MainServiceImpl struct {
	Repository repository.MainRepository
	DB         *sql.DB
}

func NewMainService(db *sql.DB, repository repository.MainRepository) MainService {
	return &MainServiceImpl{
		Repository: repository,
		DB:         db,
	}
}

func (service *MainServiceImpl) Create(ctx context.Context, data domain.MainDomain) web.MainModelResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	res, err := service.Repository.Create(ctx, tx, data)
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx, err)
	return res
}

func (service *MainServiceImpl) GetAll(ctx context.Context) []web.MainModelResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	defer func() {
		helper.CommitOrRollback(tx, err)
	}()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	res := service.Repository.GetAll(ctx, tx)
	return res
}

func (service *MainServiceImpl) GetByUsername(ctx context.Context, username string) []web.MainModelResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer func() {
		helper.CommitOrRollback(tx, err)
	}()
	res := service.Repository.GetByUsername(ctx, tx, username)
	return res
}

func (service *MainServiceImpl) GetById(ctx context.Context, id string) (web.MainModelResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer func() {
		helper.CommitOrRollback(tx, err)
	}()
	res, err := service.Repository.GetById(ctx, tx, id)
	if err != nil {
		return web.MainModelResponse{}, err
	}
	return res, nil
}

func (service *MainServiceImpl) UpdateById(ctx context.Context, data domain.MainDomain) (web.MainModelResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer func() {
		helper.CommitOrRollback(tx, err)
	}()
	res, err := service.Repository.UpdateById(ctx, tx, data)
	if err != nil {
		return web.MainModelResponse{}, err
	}
	return res, nil
}

func (service *MainServiceImpl) Delete(ctx context.Context, id string) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	res := service.Repository.Delete(ctx, tx, id)
	defer helper.CommitOrRollback(tx, err)
	return res
}
