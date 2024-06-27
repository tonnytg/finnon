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
		t.Errorf("Outcome validation failed: %v", err)
	}

	invalidOutcome := domain.Outcome{
		Description: "",
	}

	if err := invalidOutcome.Validate(); err == nil {
		t.Errorf("Expected validation error for invalid Outcome, but got none")
	}
}

func TestOutcomeRepository(t *testing.T) {
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

	// Test updating the outcome
	outcome.Description = "updated description"
	err = service.UpdateOutcome(outcome)
	if err != nil {
		t.Errorf("Error updating outcome: %v", err)
	}

	// Test finding the updated outcome
	updatedOutcome, err := service.FindOutcomeByID(outcome.ID)
	if err != nil {
		t.Errorf("Error finding outcome: %v", err)
	}
	if updatedOutcome.Description != "updated description" {
		t.Errorf("Expected description to be 'updated description', but got %v", updatedOutcome.Description)
	}

	// Test deleting the outcome
	err = service.DeleteOutcome(outcome.ID)
	if err != nil {
		t.Errorf("Error deleting outcome: %v", err)
	}

	// Test finding the deleted outcome
	_, err = service.FindOutcomeByID(outcome.ID)
	if err == nil {
		t.Errorf("Expected error when finding deleted outcome, but got none")
	}
}
