package person

import (
	"context"
	"time"
)

type Domain struct {
	Name                      string
	Gender                    string
	Desa                      string
	Kelompok                  string
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
}

type Usecase interface {
	CreatePerson(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	CreatePerson(ctx context.Context, domain *Domain) (Domain, error)
}
