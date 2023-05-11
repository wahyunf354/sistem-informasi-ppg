package person

import (
	"context"
	"time"
)

type Domain struct {
	ID                        string
	Name                      string
	Gender                    string
	DesaId                    string
	KelompokId                string
	Birth                     time.Time
	LastEduction              string
	CurrentEducation          string
	ParentName                string
	NoHp                      string
	IsPPG                     bool
	IsMubalegSetempat         bool
	IsMubalegTugasan          bool
	Dapukan                   string
	StatusPernikahan          string
	KeterampilanKemandirian   string
	Hobby                     string
	CountHadistHimpunanKhatam int
	KutubusitahKhatam         []string
	BloodGroup                string
	UrlProfile                string
	CreatedAt                 time.Time
	UpdateAt                  time.Time

	Desa     Desa
	Kelompok Kelompok
}

type Desa struct {
	Name string
}

type Kelompok struct {
	Name string
}

type Usecase interface {
	CreatePerson(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	CreatePerson(ctx context.Context, domain *Domain) (Domain, error)
}
