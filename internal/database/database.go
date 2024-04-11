package database

import (
	"database/sql"
	"finnon/internal/domain"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	conn, err := sql.Open("sqlite3", "finnon.db")
	if err != nil {
		log.Println(err)
	}

	// Create table if not exist
	createTableStatement := `
CREATE TABLE IF NOT EXISTS incomes (
    id INTEGER PRIMARY KEY,
    description TEXT,
    amount REAL,
    source TEXT,
    provider TEXT,
    payment_date DATE,
    type TEXT,
    created_at DATETIME
);

CREATE TRIGGER IF NOT EXISTS set_created_at
AFTER INSERT ON incomes
FOR EACH ROW
BEGIN
    UPDATE incomes SET created_at = DATETIME('now') WHERE id = NEW.id;
END;
    `
	_, err = conn.Exec(createTableStatement)
	if err != nil {
		log.Println(err)
	}
}

func InsertIncome(income domain.Income) {
	conn, err := sql.Open("sqlite3", "finnon.db")
	if err != nil {
		log.Println(err)
	}
	stmt, err := conn.Prepare("INSERT INTO incomes (description, amount, source, provider, payment_date, type) VALUES (?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(income.Description,
		income.Amount,
		income.Source,
		income.Provider,
		income.PaymentDate,
		income.Type)

	if err != nil {
		log.Println(err)
	}
	log.Println("New income inserted")
}

func SelectIncomes() {

	conn, err := sql.Open("sqlite3", "finnon.db")
	if err != nil {
		log.Println(err)
	}
	rows, err := conn.Query("SELECT * FROM incomes")
	if err != nil {
		log.Println("error to get information from clients")
		return
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
			&income.Type,
			&income.CreatedAt)
		if err != nil {
			log.Println("panic:", err)
		}
		log.Println("Income:", income)
	}
	if err = rows.Err(); err != nil {
		log.Println("error to read rows:", err)
	}
	log.Println("finish loop in db")
}
