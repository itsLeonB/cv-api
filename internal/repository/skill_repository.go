package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/itsLeonB/cv-api/internal/apperror"
	"github.com/itsLeonB/cv-api/internal/entity"
	"github.com/jackc/pgx/v5"
)

type SkillRepository interface {
	InsertCategory(ctx context.Context, category *entity.SkillCategory) error
	SelectCategoryByName(ctx context.Context, name string) (*entity.SkillCategory, error)
	SelectAllCategories(ctx context.Context) ([]*entity.SkillCategory, error)
	SelectCategoryByID(ctx context.Context, id int) (*entity.SkillCategory, error)
}

type skillRepository struct {
	structName string
	conn       *pgx.Conn
}

func NewSkillRepository(conn *pgx.Conn) SkillRepository {
	return &skillRepository{"skillRepository", conn}
}

func (r *skillRepository) InsertCategory(ctx context.Context, category *entity.SkillCategory) error {
	methodName := "InsertCategory()"
	return runInTx(r.conn, ctx, func(tx pgx.Tx) error {
		sql := `
			INSERT INTO skill_categories (name)
			VALUES ($1)
			RETURNING id, created_at, updated_at
		`

		err := tx.QueryRow(ctx, sql, category.Name).Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return apperror.NewAppError(
				err, r.structName, methodName,
				fmt.Sprintf("tx.QueryRow(name: %s).Scan()", category.Name),
			)
		}

		return nil
	})
}

func (r *skillRepository) SelectCategoryByName(ctx context.Context, name string) (*entity.SkillCategory, error) {
	methodName := fmt.Sprintf("SelectCategoryByName(name: %s)", name)
	sql := `
		SELECT id, name, created_at, updated_at, deleted_at
		FROM skill_categories
		WHERE name = $1
	`

	category := new(entity.SkillCategory)
	err := r.conn.QueryRow(ctx, sql, name).Scan(
		&category.ID, &category.Name,
		&category.CreatedAt, &category.UpdatedAt, &category.DeletedAt,
	)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, apperror.NewAppError(
			err, r.structName, methodName,
			"r.conn.QueryRow().Scan()",
		)
	}
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	return category, nil
}

func (r *skillRepository) SelectAllCategories(ctx context.Context) ([]*entity.SkillCategory, error) {
	methodName := "SelectAllCategories()"
	sql := `
		SELECT id, name, created_at, updated_at, deleted_at
		FROM skill_categories
	`

	rows, err := r.conn.Query(ctx, sql)
	if err != nil {
		return nil, apperror.NewAppError(
			err, r.structName, methodName,
			"r.conn.Query()",
		)
	}

	categories := []*entity.SkillCategory{}
	for rows.Next() {
		category := new(entity.SkillCategory)
		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt, &category.DeletedAt)
		if err != nil {
			return nil, apperror.NewAppError(
				err, r.structName, methodName,
				"rows.Scan()",
			)
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (r *skillRepository) SelectCategoryByID(ctx context.Context, id int) (*entity.SkillCategory, error) {
	methodName := fmt.Sprintf("SelectCategoryByID(id: %d)", id)
	sql := `
		SELECT id, name, created_at, updated_at, deleted_at
		FROM skill_categories
		WHERE id = $1
	`

	category := new(entity.SkillCategory)
	err := r.conn.QueryRow(ctx, sql, id).Scan(
		&category.ID, &category.Name,
		&category.CreatedAt, &category.UpdatedAt, &category.DeletedAt,
	)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, apperror.NewAppError(
			err, r.structName, methodName,
			"r.conn.QueryRow().Scan()",
		)
	}
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	return category, nil
}
