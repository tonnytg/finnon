package domain_test

import (
	"finnon/internal/domain"
	"finnon/internal/infra/database"
	"testing"
)

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

		})
	}
}

func TestIncomeService(t *testing.T) {

	repo := database.NewIncomeRepositorySQLiteMemory()
	service := domain.NewIncomeService(repo)
	_, err := service.CreateIncome("teste", 1001, "testeSource", "testProvider", "testePaymentDate", "testType")
	if err != nil {
		t.Errorf("cannot create income")
	}

}
