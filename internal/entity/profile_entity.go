package entity

import (
	"time"
)

type Profile struct {
	ID           int
	FullName     string
	Nickname     string
	Occupation   string
	Location     string
	ShortSummary string
	Summary      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}