package database

import (
	"database/sql"
	"finnon/internal/domain"
	"fmt"
	"log"
)

type IncomeRepositorySQLite struct{}

func NewIncomeRepositorySQLite() *IncomeRepositorySQLite {
	return &IncomeRepositorySQLite{}
}

func (r *IncomeRepositorySQLite) Save(income *domain.Income) error {

	db, err := sql.Open("sqlite3", "finnon.db")
	if err != nil {
		return fmt.Errorf("error to save in database sqlite3")
	}

	stmt, err := db.Prepare(
		"INSERT INTO incomes (description, amount, source, payment_date, provider, type, created_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("error to create stmt")
	}

	rows, err := stmt.Query(
		income.Description,
		income.Amount,
		income.Source,
		income.PaymentDate,
		income.Provider,
		income.Type,
		income.CreatedAt)
	if err != nil {
		log.Println("error to execute stmt")
	}
	for rows.Next() {
		log.Println(rows)
	}

	log.Println("value saved in SQLite3")

	return nil
}

type OutcomeRepositorySQLite struct{}

func NewOutcomeRepositorySQLite() *OutcomeRepositorySQLite {
	return &OutcomeRepositorySQLite{}
}

func (r *OutcomeRepositorySQLite) Save(outcome *domain.Outcome) error {

	db, err := sql.Open("sqlite3", "finnon.db")
	if err != nil {
		return fmt.Errorf("error to save in database sqlite3")
	}

	stmt, err := db.Prepare(
		"INSERT INTO outcomes (description, amount, tax, isPayed, companyName, installment, totalInstallment, paidAt, category, typeOutcome, repeat, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("error to create stmt")
	}

	rows, err := stmt.Query(
		outcome.Description,
		outcome.Amount,
		outcome.Tax,
		outcome.IsPayed,
		outcome.CompanyName,
		outcome.Installment,
		outcome.TotalInstallment,
		outcome.PaidAt,
		outcome.Category,
		outcome.TypeOutcome,
		outcome.Repeat,
		outcome.CreatedAt)
	if err != nil {
		log.Println("error to execute stmt")
	}
	for rows.Next() {
		log.Println(rows)
	}

	log.Println("value saved in SQLite3")

	return nil
}
