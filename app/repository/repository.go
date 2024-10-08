package repository

import (
	"context"
	"database/sql"

	"github.com/itsLeonB/cv-api/appcontext"
	"github.com/itsLeonB/cv-api/entity"
)

type Repository interface {
	GetShortSummary(context.Context) (*entity.Profile, error)
}

type repositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repositoryImpl {
	return &repositoryImpl{db}
}

func (r *repositoryImpl) GetShortSummary(ctx context.Context) (*entity.Profile, error) {
	query := `
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

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	id := ctx.Value(appcontext.KeyProfileID)
	profile := new(entity.Profile)

	err = stmt.QueryRowContext(ctx, id).Scan(
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
