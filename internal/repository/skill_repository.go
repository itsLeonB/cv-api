package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/itsLeonB/cv-api/internal/apperror"
	"github.com/itsLeonB/cv-api/internal/delivery/apphttp/httperror"
	"github.com/itsLeonB/cv-api/internal/entity"
	"github.com/jackc/pgx/v5"
)

type SkillRepository interface {
	InsertCategory(ctx context.Context, category *entity.SkillCategory) error
	SelectCategoryByName(ctx context.Context, name string) (*entity.SkillCategory, error)
	SelectAllCategories(ctx context.Context) ([]*entity.SkillCategory, error)
	SelectCategoryByID(ctx context.Context, id int) (*entity.SkillCategory, error)
	Insert(ctx context.Context, skill *entity.Skill) error
	SelectAll(ctx context.Context) ([]*entity.Skill, error)
	SelectByID(ctx context.Context, id int) (*entity.Skill, error)
	Update(ctx context.Context, skill *entity.Skill) error
	DeleteByID(ctx context.Context, id int) error
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

func (r *skillRepository) Insert(ctx context.Context, skill *entity.Skill) error {
	methodName := "Insert()"
	return runInTx(r.conn, ctx, func(tx pgx.Tx) error {
		sql := `
			INSERT INTO skills (profile_id, category_id, name, description)
			VALUES ($1, $2, $3, $4)
			RETURNING id, created_at, updated_at
		`

		err := tx.QueryRow(ctx, sql,
			skill.ProfileID, skill.CategoryID, skill.Name, skill.Description,
		).Scan(&skill.ID, &skill.CreatedAt, &skill.UpdatedAt)
		if err != nil {
			return apperror.NewAppError(
				err, r.structName, methodName,
				"tx.QueryRow().Scan()",
			)
		}

		return nil
	})
}

func (r *skillRepository) SelectAll(ctx context.Context) ([]*entity.Skill, error) {
	methodName := "SelectAll()"
	sql := `
		SELECT id, profile_id, category_id, name, description, created_at, updated_at, deleted_at
		FROM skills
	`

	rows, err := r.conn.Query(ctx, sql)
	if err != nil {
		return nil, apperror.NewAppError(
			err, r.structName, methodName,
			"r.conn.Query()",
		)
	}

	skills := []*entity.Skill{}
	for rows.Next() {
		skill := new(entity.Skill)
		err = rows.Scan(
			&skill.ID,
			&skill.ProfileID,
			&skill.CategoryID,
			&skill.Name,
			&skill.Description,
			&skill.CreatedAt,
			&skill.UpdatedAt,
			&skill.DeletedAt,
		)
		if err != nil {
			return nil, apperror.NewAppError(
				err, r.structName, methodName,
				"rows.Scan()",
			)
		}

		skills = append(skills, skill)
	}

	return skills, nil
}

func (r *skillRepository) SelectByID(ctx context.Context, id int) (*entity.Skill, error) {
	methodName := "SelectAll()"
	sql := `
		SELECT 
			skill.id, 
			skill.profile_id, 
			skill.category_id, 
			skill.name, 
			skill.description, 
			skill.created_at, 
			skill.updated_at, 
			skill.deleted_at,
			category.id,
			category.name,
			category.created_at,
			category.updated_at,
			category.deleted_at
		FROM skills skill
		JOIN skill_categories category ON skill.category_id = category.id
		WHERE skill.id = $1
	`

	skill := new(entity.Skill)
	category := new(entity.SkillCategory)
	err := r.conn.QueryRow(ctx, sql, id).Scan(
		&skill.ID,
		&skill.ProfileID,
		&skill.CategoryID,
		&skill.Name,
		&skill.Description,
		&skill.CreatedAt,
		&skill.UpdatedAt,
		&skill.DeletedAt,
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.DeletedAt,
	)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, apperror.NewAppError(
			err, r.structName, methodName,
			"r.conn.Query().Scan()",
		)
	}
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	skill.Category = category

	return skill, nil
}

func (r *skillRepository) Update(ctx context.Context, skill *entity.Skill) error {
	methodName := "Update()"
	return runInTx(r.conn, ctx, func(tx pgx.Tx) error {
		sql := `
			UPDATE skills
			SET category_id = $1, name = $2, description = $3, updated_at = NOW()
			WHERE id = $4
			RETURNING updated_at
		`

		err := tx.QueryRow(ctx, sql,
			skill.CategoryID, skill.Name, skill.Description, skill.ID,
		).Scan(&skill.UpdatedAt)
		if err != nil {
			return apperror.NewAppError(
				err, r.structName, methodName,
				"tx.QueryRow().Scan()",
			)
		}

		return nil
	})
}

func (r *skillRepository) DeleteByID(ctx context.Context, id int) error {
	methodName := fmt.Sprintf("DeleteByID(id: %d)", id)
	return runInTx(r.conn, ctx, func(tx pgx.Tx) error {
		sql := `
			UPDATE skills
			SET deleted_at = NOW()
			WHERE id = $1
		`

		cmd, err := tx.Exec(ctx, sql, id)
		if err != nil {
			return apperror.NewAppError(
				err, r.structName, methodName, "tx.Exec()",
			)
		}
		if cmd.RowsAffected() != 1 {
			return apperror.NewAppError(
				httperror.InternalServerError(),
				r.structName, methodName, "cmd.RowsAffected() != 1",
			)
		}

		return nil
	})
}
