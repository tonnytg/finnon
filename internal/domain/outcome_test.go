package domain_test

import (
	"finnon/internal/domain"
	"finnon/internal/infra/database"
	"testing"
)

func TestOutcomeValidation(t *testing.T) {
	validOutcome := domain.Outcome{
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

	if err := validOutcome.Validate(); err != nil {
		t.Errorf("Valid outcome validation failed: %v", err)
	}

	invalidOutcome := domain.Outcome{
		Description: "",
		Amount:      0,
	}

	if err := invalidOutcome.Validate(); err == nil {
		t.Errorf("Expected validation error for invalid outcome, but got none")
	}
}

func TestCreateUpdateDeleteOutcome(t *testing.T) {
	repo := database.NewOutcomeRepositorySQLiteMemory()
	service := domain.NewOutcomeService(repo)

	outcome, err := service.CreateOutcome("test outcomeRepository",
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
		t.Errorf("Error in service to create outcome: %v", err)
	}

	t.Run("Update Outcome", func(t *testing.T) {
		outcome.Description = "updated description"
		err = service.UpdateOutcome(outcome)
		if err != nil {
			t.Errorf("Error updating outcome: %v", err)
		}

		updatedOutcome, err := service.FindOutcomeByID(outcome.ID)
		if err != nil {
			t.Errorf("Error finding outcome: %v", err)
		}
		if updatedOutcome.Description != "updated description" {
			t.Errorf("Expected description to be 'updated description', but got %v", updatedOutcome.Description)
		}
	})

	t.Run("Delete Outcome", func(t *testing.T) {
		err = service.DeleteOutcome(outcome.ID)
		if err != nil {
			t.Errorf("Error deleting outcome: %v", err)
		}

		_, err = service.FindOutcomeByID(outcome.ID)
		if err == nil {
			t.Errorf("Expected error when finding deleted outcome, but got none")
		}
	})
}
