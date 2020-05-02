package model

import (
	"time"
)

type Template struct {
	ID        uint       `gorm:"primary_key" json:"id" example:"1"`
	CreatedAt time.Time  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at" example:""`
	UpdatedAt time.Time  `json:"updated_at" example:""`
	DeletedAt *time.Time `json:"deleted_at" example:""`
}
