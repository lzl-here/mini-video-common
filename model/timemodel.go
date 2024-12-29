package model

type TimeModel struct {
	CreatedAt string `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at,omitempty" gorm:"column:updated_at"`
	DeletedAt string `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}
