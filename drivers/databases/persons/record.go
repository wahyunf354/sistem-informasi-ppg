package persons

import (
	"gorm.io/gorm"
	"sistem-informasi-ppg/business/person"
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

func FromDomain(domain *person.Domain) *Person {
	return &Person{
		Name:                      domain.Name,
		Gender:                    domain.Gender,
		DesaId:                    domain.DesaId,
		KelompokId:                domain.KelompokId,
		Birth:                     domain.Birth,
		LastEducation:             domain.LastEduction,
		CurrentEducation:          domain.CurrentEducation,
		ParentName:                domain.ParentName,
		NoHp:                      domain.NoHp,
		IsPPG:                     domain.IsPPG,
		IsMubalegSetempat:         domain.IsMubalegSetempat,
		IsMubalegTugasan:          domain.IsMubalegTugasan,
		Dapukan:                   domain.Dapukan,
		StatusPernikahan:          domain.StatusPernikahan,
		KeterampilanKemandirian:   domain.KeterampilanKemandirian,
		Hobby:                     domain.Hobby,
		CountHadistHimpunanKhatam: domain.CountHadistHimpunanKhatam,
		KutubusitahKhatam:         domain.KutubusitahKhatam,
		BloodGroup:                domain.BloodGroup,
		UrlProfile:                domain.UrlProfile,
	}
}

func (p *Person) ToDomain() person.Domain {
	return person.Domain{
		ID:                        p.ID,
		Name:                      p.Name,
		Gender:                    p.Gender,
		DesaId:                    p.DesaId,
		KelompokId:                p.KelompokId,
		Birth:                     p.Birth,
		LastEduction:              p.LastEducation,
		CurrentEducation:          p.CurrentEducation,
		ParentName:                p.ParentName,
		NoHp:                      p.NoHp,
		IsPPG:                     p.IsPPG,
		IsMubalegSetempat:         p.IsMubalegSetempat,
		IsMubalegTugasan:          p.IsMubalegTugasan,
		Dapukan:                   p.Dapukan,
		StatusPernikahan:          p.StatusPernikahan,
		KeterampilanKemandirian:   p.KeterampilanKemandirian,
		Hobby:                     p.Hobby,
		CountHadistHimpunanKhatam: p.CountHadistHimpunanKhatam,
		KutubusitahKhatam:         p.KutubusitahKhatam,
		BloodGroup:                p.BloodGroup,
		UrlProfile:                p.UrlProfile,
		CreatedAt:                 p.CreatedAt,
		UpdateAt:                  p.UpdateAt,
		Desa: person.Desa{
			Name: p.Desa.Name,
		},
		Kelompok: person.Kelompok{
			Name: p.Kelompok.Name,
		},
	}
}
