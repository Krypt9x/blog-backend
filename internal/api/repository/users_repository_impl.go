package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Krypt9x/blog-backend/internal/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUsersRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) error {
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)"
	_, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) error {
	query := "SELECT username, email, password FROM users WHERE username=$1 AND email=$2 AND password=$3"
	rows, err := tx.QueryContext(ctx, query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	if rows.Next() {
		return nil
	}
	return errors.New("not found")
}
