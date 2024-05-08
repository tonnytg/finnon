package domain

import "log"

type OutcomeRepositoryInterface interface {
	Save(*Outcome) error
}

type OutcomeRepository struct{}

func NewOutcomeRepository() *OutcomeRepository {
	return &OutcomeRepository{}
}

func (r *OutcomeRepository) Save(outcome *Outcome) error {
	log.Println("outcome:", outcome, "pass in OutcomeRepository.")
	return nil
}
