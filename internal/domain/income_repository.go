package domain

import "log"

type IncomeRepositoryInterface interface {
	Save(*Income) error
}

type IncomeRepository struct{}

func NewIncomeRepository() *IncomeRepository {
	return &IncomeRepository{}
}

func (r *IncomeRepository) Save(income *Income) error {
	log.Println("income:", income, "pass in IncomeRepository.")
	return nil
}
