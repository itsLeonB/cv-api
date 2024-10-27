package repository

import (
	"context"
	"fmt"

	"github.com/itsLeonB/cv-api/internal/apperror"
	"github.com/itsLeonB/cv-api/internal/entity"
	"github.com/jackc/pgx/v5"
)

type ProfileRepository interface {
	GetShortSummaryByID(ctx context.Context, id int) (*entity.Profile, error)
	GetSummaryByID(ctx context.Context, id int) (*entity.Profile, error)
}

type profileRepository struct {
	structName string
	conn       *pgx.Conn
}

func NewProfileRepository(conn *pgx.Conn) *profileRepository {
	return &profileRepository{"profileRepository", conn}
}

func (r *profileRepository) GetShortSummaryByID(ctx context.Context, id int) (*entity.Profile, error) {
	methodName := fmt.Sprintf("GetShortSummaryByID(id: %d)", id)
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
		return nil, apperror.NewAppError(
			err, r.structName, methodName,
			"r.conn.QueryRow().Scan()",
		)
	}

	return profile, nil
}

func (r *profileRepository) GetSummaryByID(ctx context.Context, id int) (*entity.Profile, error) {
	methodName := fmt.Sprintf("GetSummaryByID(id: %d)", id)
	sql := `
		SELECT full_name, summary
		FROM profiles
		WHERE deleted_at IS NULL AND id = $1
	`

	profile := new(entity.Profile)
	err := r.conn.QueryRow(ctx, sql, id).Scan(
		&profile.FullName,
		&profile.Summary,
	)
	if err != nil {
		return nil, apperror.NewAppError(
			err, r.structName, methodName,
			"r.conn.QueryRow().Scan()",
		)
	}

	return profile, nil
}
