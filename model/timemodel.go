package model

import "time"

type TimeModel struct {
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}
