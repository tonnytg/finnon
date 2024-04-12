package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

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

func init() {

	conn := NewDB(nil)

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
	_, err := conn.db.Exec(createTableStatement)
	if err != nil {
		log.Println(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("error to close connection in init")
		}
	}(conn.db)
}
