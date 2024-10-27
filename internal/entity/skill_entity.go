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
