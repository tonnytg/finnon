package domain

import (
	"fmt"
)

type IncomeService struct {
	repo IncomeRepositoryInterface
}

func NewIncomeService(repo IncomeRepositoryInterface) *IncomeService {
	return &IncomeService{repo}
}

func (s *IncomeService) CreateIncome(description string, amount float64, source string, provider string, paymentDate string, incomeType string) (*Income, error) {

	if description == "" || amount < 1 || source == "" || provider == "" || paymentDate == "" || incomeType == "" {
		return nil, fmt.Errorf("errror to save in repository because values cannot be empty")
	}

	income := NewIncome(description, amount, source, provider, paymentDate, incomeType)
	err := income.Validate()
	if err != nil {
		return nil, fmt.Errorf("sorry Income has error")
	}

	err = s.repo.Save(income)
	if err != nil {
		return nil, fmt.Errorf("errror to save in repository")
	}

	return income, nil
}
