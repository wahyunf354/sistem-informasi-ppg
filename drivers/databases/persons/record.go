package persons

import (
	"gorm.io/gorm"
	"sistem-informasi-ppg/drivers/databases/desas"
	"sistem-informasi-ppg/drivers/databases/kelompoks"
	"time"
)

type Person struct {
	ID                        string             `gorm:"primaryKey;unique;not null"`
	Name                      string             `gorm:"not null"`
	Gender                    string             `gorm:"not null;type:ENUM('Laki-laki', 'Perempuan')"`
	DesaId                    string             `gorm:"not null"`
	KelompokId                string             `gorm:"not null"`
	Desa                      desas.Desa         `gorm:"foreignKey:DesaId"`
	Kelompok                  kelompoks.Kelompok `gorm:"foreignKey:KelompokId"`
	Birth                     time.Time
	LastEducation             string `gorm:"not null"`
	CurrentEducation          string
	ParentName                string `gorm:"not null"`
	NoHp                      string
	IsPPG                     bool `gorm:"not null"`
	IsMubalegSetempat         bool `gorm:"not null"`
	IsMubalegTugasan          bool `gorm:"not null"`
	Dapukan                   string
	StatusPernikahan          string `gorm:"not null;type:ENUM('Sudah Menikah', 'Belum Menikah', 'Janda', 'Duda')"`
	KeterampilanKemandirian   string
	Hobby                     string
	CountHadistHimpunanKhatam int
	KutubusitahKhatam         []string
	BloodGroup                string `gorm:"type:ENUM('A', 'B', 'AB', '0', '-A', '-B', '-AB', '-0')"`
	UrlProfile                string
	CreatedAt                 time.Time
	UpdateAt                  time.Time
	DeletedAt                 gorm.DeletedAt
}
