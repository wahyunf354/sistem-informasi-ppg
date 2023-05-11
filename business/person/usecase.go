package person

import (
	"context"
	"sistem-informasi-ppg/controllers"
	"time"
)

type PersonUsecase struct {
	PersonRepository Repository
	ContextTimeOut   time.Duration
}

func (p PersonUsecase) CreatePerson(ctx context.Context, domain *Domain) (Domain, error) {

	if domain.Name == "" {
		return Domain{}, controllers.EMPTY_NAME
	}

	result, err := p.PersonRepository.CreatePerson(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}



