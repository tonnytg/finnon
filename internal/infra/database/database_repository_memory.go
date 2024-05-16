package database

import (
	"database/sql"
	"finnon/internal/domain"
	"fmt"
	"log"
)

type IncomeRepositorySQLiteMemory struct{}

func NewIncomeRepositorySQLiteMemory() *IncomeRepositorySQLiteMemory {
	return &IncomeRepositorySQLiteMemory{}
}

func (r *IncomeRepositorySQLiteMemory) Save(income *domain.Income) error {

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return fmt.Errorf("error to save in database sqlite3")
	}

	// Create table if not exist

	_, err = db.Exec(CreateTableStatementIncomes)
	if err != nil {
		log.Println(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("error to close connection in init memory sqlite3")
		}
	}(db)
	stmt, err := db.Prepare("INSERT INTO incomes (amount, description) VALUES (?, ?)")
	if err != nil {
		log.Println("error to create smtp")
	}

	rows, err := stmt.Query(income.Amount, income.Description)
	if err != nil {
		log.Println("error to execute stmt")
	}
	for rows.Next() {
		log.Println(rows)
	}

	log.Println("value saved in SQLite3")

	return nil
}

type OutcomeRepositorySQLiteMemory struct{}

func NewOutcomeRepositorySQLiteMemory() *OutcomeRepositorySQLiteMemory {
	return &OutcomeRepositorySQLiteMemory{}
}

func (r *OutcomeRepositorySQLiteMemory) Save(outcome *domain.Outcome) error {

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return fmt.Errorf("error to save in database sqlite3")
	}

	// Create table if not exist

	_, err = db.Exec(CreateTableStatementOutcomes)
	if err != nil {
		log.Println(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("error to close connection in init memory sqlite3")
		}
	}(db)

	stmt, err := db.Prepare("INSERT INTO outcomes (description, amount) VALUES (?, ?)")
	if err != nil {
		log.Println("error to create smtp")
	}

	rows, err := stmt.Query(outcome.Description, outcome.Amount)
	if err != nil {
		log.Println("error to execute stmt")
	}
	for rows.Next() {
		log.Println(rows)
	}

	log.Println("outcome saved in SQLite3 Memory")

	return nil
}
