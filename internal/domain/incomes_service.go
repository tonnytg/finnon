package domain

import "fmt"

type IncomeService struct {
	repo IncomeRepositoryInterface
}

func NewIncomeService(repo IncomeRepositoryInterface) *IncomeService {
	return &IncomeService{repo}
}

func (s *IncomeService) CreateIncome(amount float64, description string) (*Income, error) {

	income := &Income{}
	income.SetAmount(amount)
	income.SetDescription(description)
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
