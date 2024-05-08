package domain

import (
	"fmt"
)

type OutcomeService struct {
	repo OutcomeRepositoryInterface
}

func NewOutcomeService(repo OutcomeRepositoryInterface) *OutcomeService {
	return &OutcomeService{repo}
}

func (s *OutcomeService) CreateOutcome(
	description string,
	amount float64,
	tax float64,
	isPayed bool,
	companyName string,
	installment uint,
	totalInstallment uint,
	paidAt string,
	category string,
	typeOutcome string,
	repeat bool) (*Outcome, error) {

	outcome := NewOutcome(description, amount, tax, isPayed, companyName, installment, totalInstallment, paidAt, category, typeOutcome, repeat)
	err := outcome.Validate()
	if err != nil {
		return nil, fmt.Errorf("sorry outcome has error")
	}

	err = s.repo.Save(outcome)
	if err != nil {
		return nil, fmt.Errorf("errror to save outcome in repository")
	}

	return outcome, nil
}
