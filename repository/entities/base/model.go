package base

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID != "" {
		return
	}
	m.ID = uuid.NewString()
	return
}
