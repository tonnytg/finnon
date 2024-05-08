package domain

import (
	"errors"
	"time"
)

type Outcome struct {
	ID               int     `json:"id"`
	Description      string  `json:"description"`
	Amount           float64 `json:"amount"`
	Tax              float64 `json:"tax"`
	IsPayed          bool    `json:"is_payed"`
	CompanyName      string  `json:"company_name"`
	Installment      uint    `json:"installment"`
	TotalInstallment uint    `json:"total_installment"`
	PaidAt           string  `json:"paid_at"`
	Category         string  `json:"category"`
	TypeOutcome      string  `json:"type"`
	Repeat           bool    `json:"repeat"`
	CreatedAt        string  `json:"created_at"`
}

func NewOutcome(
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
	repeat bool) *Outcome {
	return &Outcome{
		Description:      description,
		Amount:           amount,
		Tax:              tax,
		IsPayed:          isPayed,
		CompanyName:      companyName,
		Installment:      installment,
		TotalInstallment: totalInstallment,
		PaidAt:           paidAt,
		Category:         category,
		TypeOutcome:      typeOutcome,
		Repeat:           repeat,
		CreatedAt:        time.Now().Format("2006-01-02 15:04:05"),
	}
}

func (o *Outcome) GetID() int {
	return o.ID
}

func (o *Outcome) SetID(id int) {
	o.ID = id
}

func (o *Outcome) GetDescription() string {
	return o.Description
}

func (o *Outcome) SetDescription(description string) {
	o.Description = description
}

func (o *Outcome) GetAmount() float64 {
	return o.Amount
}

func (o *Outcome) SetAmount(amount float64) {
	o.Amount = amount
}

func (o *Outcome) GetTax() float64 {
	return o.Tax
}

func (o *Outcome) SetTax(tax float64) {
	o.Tax = tax
}

func (o *Outcome) GetIsPayed() bool {
	return o.IsPayed
}

func (o *Outcome) SetIsPayed(isPayed bool) {
	o.IsPayed = isPayed
}

func (o *Outcome) GetCompanyName() string {
	return o.CompanyName
}

func (o *Outcome) SetCompanyName(companyName string) {
	o.CompanyName = companyName
}

func (o *Outcome) GetInstallment() uint {
	return o.Installment
}

func (o *Outcome) SetInstallment(installment uint) {
	o.Installment = installment
}

func (o *Outcome) GetTotalInstallment() uint {
	return o.TotalInstallment
}

func (o *Outcome) SetTotalInstallment(totalInstallment uint) {
	o.TotalInstallment = totalInstallment
}

func (o *Outcome) GetPaidAt() string {
	return o.PaidAt
}

func (o *Outcome) SetPaidAt(paidAt string) {
	o.PaidAt = paidAt
}

func (o *Outcome) GetCategory() string {
	return o.Category
}

func (o *Outcome) SetCategory(category string) {
	o.Category = category
}

func (o *Outcome) GetType() string {
	return o.TypeOutcome
}

func (o *Outcome) SetType(typeOutcome string) {
	o.TypeOutcome = typeOutcome
}

func (o *Outcome) GetRepeat() bool {
	return o.Repeat
}

func (o *Outcome) SetRepeat(repeat bool) {
	o.Repeat = repeat
}

func (o *Outcome) GetCreatedAt() string {
	return o.CreatedAt
}

func (o *Outcome) SetCreatedAt(createdAt string) {
	o.CreatedAt = createdAt
}

func (o *Outcome) Validate() error {
	if o.Amount == 0 || o.Description == "" || o.TotalInstallment == 0 || o.Installment == 0 || o.Category == "" {
		return errors.New("outcome cannot has empty values")
	}
	return nil
}
