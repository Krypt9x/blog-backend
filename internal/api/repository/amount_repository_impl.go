package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Krypt9x/blog-backend/pkg/helper"
)

type AmountRepositoryImpl struct {
}

func NewAmountRepository() AmountRepository {
	return &AmountRepositoryImpl{}
}

func (repository *AmountRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, id string) error {
	query := "INSERT INTO amounts(id_post, amount_comments, amount_views) VALUES($1, $2, $3)"
	_, err := tx.ExecContext(ctx, query, id, 0, 0)
	if err != nil {
		return err
	}
	return nil
}

func (repository *AmountRepositoryImpl) GetAmountCommentsById(ctx context.Context, tx *sql.Tx, id string) (uint64, error) {
	query := "SELECT amount_comments FROM amounts WHERE id_post=$1"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return 0, err
	}

	var amountComment uint64
	if rows.Next() {
		err := rows.Scan(&amountComment)
		helper.PanicIfError(err)
	}

	return amountComment, nil
}

func (repository *AmountRepositoryImpl) GetAmountViewsById(ctx context.Context, tx *sql.Tx, id string) (uint64, error) {
	query := "SELECT amount_views FROM amounts WHERE id_post=$1"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return 0, err
	}

	var amountViews uint64
	if rows.Next() {
		err := rows.Scan(&amountViews)
		helper.PanicIfError(err)
	}

	return amountViews, nil
}

func (repository *AmountRepositoryImpl) UpdateAmountCommentsById(ctx context.Context, tx *sql.Tx, id string) (uint64, error) {
	query := "UPDATE amounts SET amount_comments=amount_comments+1 WHERE id_post=$1 RETURNING amount_comments"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return 0, err
	}
	var UpdatedAmountCommments uint64
	if rows.Next() {
		err := rows.Scan(&UpdatedAmountCommments)
		if err != nil {
			return 0, err
		}

		return UpdatedAmountCommments, nil
	}
	return 0, errors.New("data not found")
}

func (repository *AmountRepositoryImpl) UpdateAmountViewsById(ctx context.Context, tx *sql.Tx, id string) (uint64, error) {
	query := "UPDATE amounts SET amount_views=amount_views+1 WHERE id_post=$1 RETURNING amount_views"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return 0, err
	}
	var UpdatedAmountViews uint64
	if rows.Next() {
		err := rows.Scan(&UpdatedAmountViews)
		if err != nil {
			return 0, err
		}

		return UpdatedAmountViews, nil
	}
	return 0, errors.New("data not found")
}

func (repository *AmountRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	query := "DELETE FROM amounts WHERE id_post=$1"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
