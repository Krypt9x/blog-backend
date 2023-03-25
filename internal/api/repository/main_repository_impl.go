package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Krypt9x/blog-backend/internal/model/domain"
	"github.com/Krypt9x/blog-backend/internal/model/web"
)

type MainRepositoryImpl struct {
}

func NewMainRepository() MainRepository {
	return &MainRepositoryImpl{}
}

func (repository *MainRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, data domain.MainDomain) (web.MainModelResponse, error) {
	query := "INSERT INTO main_table(id_post, title, username, date, trailer_content, content) VALUES ($1, $2, $3, $4, $5, $6)"
	timeNow := time.Now()
	timeNowArr := []string{strconv.Itoa(timeNow.Hour()), strconv.Itoa(timeNow.Minute()), " at ", strconv.Itoa(timeNow.Day()), strconv.Itoa(int(timeNow.Month())), strconv.Itoa(timeNow.Year())}
	contentDate := strings.Join(timeNowArr, "-")

	_, err := tx.ExecContext(ctx, query, data.Id, data.Title, data.Username, contentDate, data.TrailerContent, data.Content)
	if err != nil {
		log.Println(err)
		return web.MainModelResponse{}, err
	}

	res := web.MainModelResponse{
		Id:             data.Id,
		Title:          data.Title,
		Username:       data.Username,
		Date:           contentDate,
		TrailerContent: data.TrailerContent,
		Content:        data.Content,
	}

	return res, nil
}

func (repository *MainRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []web.MainModelResponse {
	query := "SELECT id_post, title, username, date, trailer_content, content FROM main_table ORDER BY no DESC"
	var allData []web.MainModelResponse
	var data web.MainModelResponse
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&data.Id, &data.Title, &data.Username, &data.Date, &data.TrailerContent, &data.Content)
		if err != nil {
			panic(err)
		}
		allData = append(allData, data)
	}
	return allData
}

func (repository *MainRepositoryImpl) GetByUsername(ctx context.Context, tx *sql.Tx, username string) []web.MainModelResponse {
	query := "SELECT id_post, title, username, date, trailer_content, content FROM main_table WHERE username=$1 ORDER BY no DESC"
	var data web.MainModelResponse
	var allData []web.MainModelResponse
	rows, err := tx.QueryContext(ctx, query, username)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&data.Id, &data.Title, &data.Username, &data.Date, &data.TrailerContent, &data.Content)
		if err != nil {
			panic(err)
		}

		allData = append(allData, data)
	}

	return allData
}

func (repository *MainRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, id string) (web.MainModelResponse, error) {
	query := "SELECT id_post, username, title, date, trailer_content, content FROM main_table WHERE id_post=$1"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	var data web.MainModelResponse
	if rows.Next() {
		err := rows.Scan(&data.Id, &data.Title, &data.Username, &data.Date, &data.TrailerContent, &data.Content)
		if err != nil {
			panic(err)
		}

		return data, nil
	}
	return web.MainModelResponse{}, errors.New("data not found")
}

func (repository *MainRepositoryImpl) UpdateById(ctx context.Context, tx *sql.Tx, data domain.MainDomain) (web.MainModelResponse, error) {
	query := "UPDATE main_table SET title=$1, trailer_content=$2, content=$3 WHERE id_post=$4"
	_, err := tx.ExecContext(ctx, query, data.Title, data.TrailerContent, data.Content, data.Id)
	if err != nil {
		return web.MainModelResponse{}, nil
	}

	res := web.MainModelResponse{
		Title:          data.Title,
		TrailerContent: data.TrailerContent,
		Content:        data.Content,
	}

	return res, nil
}

func (repository *MainRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	query := "DELETE FROM main_table WHERE id_post=$1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		panic(err)
	}
	return nil
}
