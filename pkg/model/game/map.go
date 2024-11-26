package game

import (
	"time"
)

type Map struct {
	Id        string `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Maps []*Map

