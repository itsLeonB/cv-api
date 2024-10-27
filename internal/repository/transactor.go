package repository

import (
	"context"

	"github.com/itsLeonB/cv-api/internal/apperror"
	"github.com/jackc/pgx/v5"
)

func runInTx(conn *pgx.Conn, ctx context.Context, fn func(tx pgx.Tx) error) error {
	methodName := "runInTx()"
	tx, err := conn.Begin(ctx)
	if err != nil {
		return apperror.NewAppError(err, "", methodName, "conn.Begin(ctx)")
	}
	defer tx.Rollback(ctx)

	err = fn(tx)
	if err != nil {
		return apperror.NewAppError(err, "", methodName, "fn(tx)")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return apperror.NewAppError(err, "", methodName, "tx.Commit(ctx)")
	}

	return nil
}
