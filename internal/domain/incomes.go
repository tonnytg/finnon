package domain

import (
	"fmt"
	"time"
)

// Income
// Amount 1.000
// PaymentDate -> 10/04/2024 10:30
// Source ->  Company Salary
// Provider -> Banks Names
// Type -> Salary, Rent, Dividend Yield
type Income struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Source      string  `json:"source"`
	PaymentDate string  `json:"payment_date"`
	Provider    string  `json:"provider"`
	TypeIncome  string  `json:"type_income"`
	CreatedAt   string  `json:"created_at"`
}

func NewIncome(description string, amount float64, source, provider, paymentDate, incomeType string) *Income {
	return &Income{
		Description: description,
		Amount:      amount,
		Source:      source,
		Provider:    provider,
		PaymentDate: paymentDate,
		TypeIncome:  incomeType,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
}

func (i *Income) GetID() int {
	return i.ID
}

func (i *Income) GetDescription() string {
	return i.Description
}

func (i *Income) GetAmount() float64 {
	return i.Amount
}

func (i *Income) GetSource() string {
	return i.Source
}

func (i *Income) GetProvider() string {
	return i.Provider
}

func (i *Income) GetPaymentDate() string {
	return i.PaymentDate
}

func (i *Income) GetType() string {
	return i.TypeIncome
}

func (i *Income) GetCreatedAt() string {
	return i.CreatedAt
}

func (i *Income) SetID(id int) {
	i.ID = id
}

func (i *Income) SetDescription(description string) {
	i.Description = description
}

func (i *Income) SetAmount(amount float64) {
	i.Amount = amount
}

func (i *Income) SetSource(source string) {
	i.Source = source
}

func (i *Income) SetProvider(provider string) {
	i.Provider = provider
}

func (i *Income) SetPaymentDate(paymentDate string) {
	i.PaymentDate = paymentDate
}

func (i *Income) SetType(incomeType string) {
	i.TypeIncome = incomeType
}

func (i *Income) Validate() error {
	if i.GetAmount() < 0 {
		return fmt.Errorf("amount cannot negative")
	}

	if i.GetDescription() == "" {
		return fmt.Errorf("description cannot be empty")
	}

	return nil
}
