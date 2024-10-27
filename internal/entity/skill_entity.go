package entity

import (
	"database/sql"
	"time"
)

type SkillCategory struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

type Skill struct {
	ID          int
	ProfileID   int
	CategoryID  int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
	Category    *SkillCategory
}
