package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/itsLeonB/cv-api/internal/apperror"
	"github.com/itsLeonB/cv-api/internal/entity"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	Insert(ctx context.Context, user *entity.User) error
	SelectByEmail(ctx context.Context, email string) (*entity.User, error)
}

type userRepository struct {
	structName string
	conn       *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) UserRepository {
	return &userRepository{"userRepository", conn}
}

func (u *userRepository) Insert(ctx context.Context, user *entity.User) error {
	methodName := "Insert()"
	return runInTx(u.conn, ctx, func(tx pgx.Tx) error {
		sql := `
			INSERT INTO users (email, password)
			VALUES ($1, $2)
			RETURNING id, created_at, updated_at
		`
		err := tx.QueryRow(ctx, sql, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return apperror.NewAppError(err, u.structName, methodName, "u.conn.QueryRow()")
		}

		return nil
	})
}

func (u *userRepository) SelectByEmail(ctx context.Context, email string) (*entity.User, error) {
	methodName := fmt.Sprintf("SelectByEmail(email: %s)", email)
	user := entity.User{Email: email}

	sql := `
		SELECT id, password, created_at, updated_at, deleted_at
		FROM users
		WHERE email = $1
	`
	err := u.conn.QueryRow(ctx, sql, email).Scan(&user.ID, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, apperror.NewAppError(err, u.structName, methodName, "u.conn.QueryRow().Scan()")
	}
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	return &user, nil
}
