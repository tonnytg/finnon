package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Db *sql.DB
}

var (
	CreateTableStatementIncomes = `
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

	CreateTableStatementOutcomes = `CREATE TABLE IF NOT EXISTS outcomes (
    id INTEGER PRIMARY KEY,
    description TEXT,
    amount REAL,
    tax REAL,
    isPayed INTEGER,
    companyName TEXT,
    installment INTEGER,
    totalInstallment INTEGER,
    paidAt DATETIME,
    category TEXT,
    typeOutcome TEXT,
    repeat TEXT,
    created_at DATETIME
);

CREATE TRIGGER IF NOT EXISTS set_created_at
AFTER INSERT ON outcomes
FOR EACH ROW
BEGIN
    UPDATE outcomes SET created_at = DATETIME('now') WHERE id = NEW.id;
END;`
)

func NewDB(db *sql.DB) *DB {

	if db == nil {
		var err error
		db, err = sql.Open("sqlite3", "finnon.db")
		if err != nil {
			log.Println("error to create connection with database sqlite3")
		}
	}
	return &DB{
		db,
	}
}
