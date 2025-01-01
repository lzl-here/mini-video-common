package model

type TimeModel struct {
	CreatedAt uint64 `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt uint64 `json:"updated_at,omitempty" gorm:"column:updated_at"`
	DeletedAt uint64 `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}
