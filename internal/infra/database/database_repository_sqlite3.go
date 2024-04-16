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
