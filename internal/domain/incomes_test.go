func TestIncome(t *testing.T) {
	tests := []struct {
		ID          int
		Description string
		Amount      float64
		Source      string
		Provider    string
		PaymentDate string
		Type        string
		CreateAt    string
	}{
		{
			ID:          1,
			Description: "Salary January",
			Amount:      1000,
			Source:      "Banco BV",
			Provider:    "Itaú",
			PaymentDate: "01/01/2024",
			Type:        "Salary",
		},
		{
			ID:          2,
			Description: "Salary February",
			Amount:      2000,
			Source:      "CLT",
			Provider:    "Banco BV",
			PaymentDate: "01/02/2024",
			Type:        "Salary",
		},
		{
			ID:          3,
			Description: "HCTR11 rent",
			Amount:      30,
			Source:      "FII",
			Provider:    "Ion",
			PaymentDate: "01/03/2024",
			Type:        "Dividend Yield",
		},
		{
			ID:          4,
			Description: "Invalid Income",
			Amount:      -10, // Amount should not be negative
			Source:      "FII",
			Provider:    "Ion",
			PaymentDate: "01/04/2024",
			Type:        "Dividend Yield",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			i := domain.NewIncome(
				tt.Description,
				tt.Amount,
				tt.Source,
				tt.Provider,
				tt.PaymentDate,
				tt.Type)
			if got := i.GetDescription(); got != tt.Description {
				t.Errorf("Income.GetDescription() = %v, want %v", got, tt.Description)
			}

			if got := i.GetAmount(); got != tt.Amount {
				t.Errorf("Income.GetAmount = %v, want %v", got, tt.Amount)
			}

			if got := i.GetPaymentDate(); got != tt.PaymentDate {
				t.Errorf("Income.GetPaymentDate = %v, want %v", got, tt.PaymentDate)
			}

			// Additional validation tests
			if tt.Amount < 0 {
				// Verify that an income with a negative amount is not valid
				t.Errorf("Income amount should not be negative")
			}
		})
	}
}

func TestIncomeService(t *testing.T) {
	repo := database.NewIncomeRepositorySQLiteMemory()
	service := domain.NewIncomeService(repo)

	// Test creating income
	createdIncome, err := service.CreateIncome("teste", 1001, "testeSource", "testProvider", "testePaymentDate", "testType")
	if err != nil {
		t.Errorf("Error creating income: %v", err)
	}

	// Test fetching income
	fetchedIncome, err := service.GetIncomeByID(createdIncome.ID)
	if err != nil {
		t.Errorf("Error fetching income: %v", err)
	}
	if fetchedIncome.Description != "teste" {
		t.Errorf("Expected fetched income description to be 'teste', got '%s'", fetchedIncome.Description)
	}

	// Test updating income
	updatedIncome := fetchedIncome
	updatedIncome.Description = "Updated Test"
	err = service.UpdateIncome(updatedIncome)
	if err != nil {
		t.Errorf("Error updating income: %v", err)
	}

	// Test deleting income
	err = service.DeleteIncome(updatedIncome.ID)
	if err != nil {
		t.Errorf("Error deleting income: %v", err)
	}
}
