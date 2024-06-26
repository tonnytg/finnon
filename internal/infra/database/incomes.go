package database

import (
	"finnon/internal/domain"
	"log"
)

func InsertIncome(db *DB, income domain.Income) {

	conn := db
	stmt, err := conn.Db.Prepare("INSERT INTO incomes (description, amount, source, provider, payment_date, type) VALUES (?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(income.Description,
		income.Amount,
		income.Source,
		income.Provider,
		income.PaymentDate,
		income.TypeIncome)

	if err != nil {
		log.Println(err)
	}
	log.Println("New income inserted")
}

func SelectIncomes(db *DB) []domain.Income {

	conn := db

	var incomes []domain.Income

	rows, err := conn.Db.Query("SELECT * FROM incomes")
	if err != nil {
		log.Println("error to get information from clients")
		return nil
	}
	defer rows.Close()

	for rows.Next() {

		var income domain.Income

		err = rows.Scan(
			&income.ID,
			&income.Description,
			&income.Amount,
			&income.Source,
			&income.Provider,
			&income.PaymentDate,
			&income.TypeIncome,
			&income.CreatedAt)
		if err != nil {
			log.Println("panic:", err)
		}
		log.Println("Income:", income)
		incomes = append(incomes, income)
	}
	if err = rows.Err(); err != nil {
		log.Println("error to read rows:", err)
	}
	return incomes
}
