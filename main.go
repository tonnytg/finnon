package main

import (
	"database/sql"
	"finnon/internal/infra/database"
	"finnon/internal/infra/webserv"
	"log"
)

func init() {

	conn := database.NewDB(nil)

	_, err := conn.Db.Exec(database.CreateTableStatementIncomes)
	if err != nil {
		log.Println(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("error to close connection in init")
		}
	}(conn.Db)

	_, err = conn.Db.Exec(database.CreateTableStatementOutcomes)
	if err != nil {
		log.Println(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("error to close connection in init")
		}
	}(conn.Db)
}

func main() {
	log.Println("Start Finnon")

	webserv.Start()
}
