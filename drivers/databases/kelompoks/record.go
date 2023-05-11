package kelompoks

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Kelompok struct {
	ID   string `gorm:"primaryKey;unique;not null"`
	Name string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (k *Kelompok) BeforeCreate(db *gorm.DB) error {
	k.ID = uuid.NewV4().String()
	k.CreatedAt = time.Now().Local()

	return nil
}

func (k *Kelompok) BeforeUpdate(db *gorm.DB) error {
	k.UpdatedAt = time.Now().Local()
	return nil
}
