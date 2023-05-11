package persons

import (
	"context"
	"gorm.io/gorm"
	"sistem-informasi-ppg/business/person"
)

type PersonRepository struct {
	Conn *gorm.DB
}

func (p *PersonRepository) CreatePerson(ctx context.Context, domain *person.Domain) (person.Domain, error) {

	rec := FromDomain(domain)
	result := p.Conn.Create(&rec)

	if result.Error != nil {
		return rec.ToDomain(), result.Error
	}

	return rec.ToDomain(), nil
}
