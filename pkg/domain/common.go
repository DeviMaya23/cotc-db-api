package domain

import (
	"time"

	"gorm.io/gorm"
)

type CommonModel struct {
	ID        int64          `json:"id" gorm:"column:id"`
	CreatedBy string         `json:"created_by,omitempty" gorm:"column:created_by"`
	UpdatedBy string         `json:"updated_by,omitempty" gorm:"column:updated_by"`
	DeletedBy *string        `json:"deleted_by,omitempty" gorm:"column:deleted_by"`
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
