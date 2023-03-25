package repository

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/Krypt9x/blog-backend/internal/model/domain"
	"github.com/Krypt9x/blog-backend/internal/model/web"
	"github.com/Krypt9x/blog-backend/pkg/helper"
)

type CommentRepositoryImpl struct {
}

func NewCommentRepository() CommentRepository {
	return &CommentRepositoryImpl{}
}

func (repository *CommentRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, data domain.Comments) error {
	query := "INSERT INTO comments(id_post, username, comment, date_comment) VALUES ($1, $2, $3, $4)"
	timeNow := time.Now()
	timeNowArr := []string{strconv.Itoa(timeNow.Hour()), strconv.Itoa(timeNow.Minute()), " at ", strconv.Itoa(timeNow.Day()), strconv.Itoa(int(timeNow.Month())), strconv.Itoa(timeNow.Year())}
	commentDateNow := strings.Join(timeNowArr, "-")
	_, err := tx.ExecContext(ctx, query, data.Id, data.Username, data.Comment, commentDateNow)
	helper.PanicIfError(err)

	return nil
}

func (repository *CommentRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, id string) ([]web.CommentResponse, error) {
	query := "SELECT id_post, username, comment, date_comment FROM comments WHERE id_post=$1"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return []web.CommentResponse{}, err
	}

	var data web.CommentResponse
	var allData []web.CommentResponse
	for rows.Next() {
		err := rows.Scan(&data.Id, &data.Username, &data.Comment, &data.DateComment)
		helper.PanicIfError(err)
		allData = append(allData, data)
	}

	return allData, nil
}

func (repository *CommentRepositoryImpl) UpdateById(ctx context.Context, tx *sql.Tx, data domain.Comments) error {
	query := "UPDATE comments SET comments=$1, date_comment=$2 WHERE id_post=$3"
	timeNow := time.Now()
	timeNowArr := []string{strconv.Itoa(timeNow.Hour()), strconv.Itoa(timeNow.Minute()), " at ", strconv.Itoa(timeNow.Day()), strconv.Itoa(int(timeNow.Month())), strconv.Itoa(timeNow.Year())}
	commentDateUpdated := strings.Join(timeNowArr, "-")
	_, err := tx.ExecContext(ctx, query, data.Comment, commentDateUpdated, data.Id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *CommentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	query := "DELETE FROM comments WHERE id_post=$1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		panic(err)
	}
	return nil
}
