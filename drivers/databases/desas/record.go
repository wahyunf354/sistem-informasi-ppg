package desas

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Desa struct {
	ID   string `gorm:"primaryKey;unique;not null"`
	Name string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (m *Desa) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewV4().String()
	m.CreatedAt = time.Now().Local()

	return nil
}

func (m *Desa) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}
