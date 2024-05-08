package domain_test

import (
	"finnon/internal/domain"
	"finnon/internal/infra/database"
	"testing"
)

func TestOutcome(t *testing.T) {

	outcome := domain.Outcome{
		Description:      "test",
		Amount:           100,
		Tax:              0,
		IsPayed:          true,
		CompanyName:      "bv",
		Installment:      1,
		TotalInstallment: 10,
		PaidAt:           "10/10/2024T10:10",
		Category:         "Restaurant",
		TypeOutcome:      "luxury",
		Repeat:           true,
		CreatedAt:        "10/10/2024T10:10",
	}

	if err := outcome.Validate(); err != nil {
		t.Errorf("cannot be empty")
	}
}

func TestOutcomeRepository(t *testing.T) {

	repo := database.NewOutcomeRepositorySQLiteMemory()
	service := domain.NewOutcomeService(repo)

	_, err := service.CreateOutcome("test outcomeRepository",
		100,
		10,
		true,
		"bv",
		1,
		10,
		"10/10/2024T10:10",
		"restaurant",
		"food",
		true)
	if err != nil {
		t.Errorf("error in service to create outcome")
	}
}
