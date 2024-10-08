package repository

import (
	"context"

	"github.com/itsLeonB/cv-api/internal/entity"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	GetShortSummaryByID(ctx context.Context, id int) (*entity.Profile, error)
}

type repository struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) *repository {
	return &repository{conn}
}

func (r *repository) GetShortSummaryByID(ctx context.Context, id int) (*entity.Profile, error) {
	sql := `
		SELECT
			id,
			nickname,
			occupation,
			location,
			short_summary
		FROM profiles
		WHERE 
			deleted_at IS NULL
			AND id = $1
	`

	profile := new(entity.Profile)
	err := r.conn.QueryRow(ctx, sql, id).Scan(
		&profile.ID,
		&profile.Nickname,
		&profile.Occupation,
		&profile.Location,
		&profile.ShortSummary,
	)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
