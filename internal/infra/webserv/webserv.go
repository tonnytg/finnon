package webserv

import (
	"log"
	"net/http"
)

func Start() {

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/incomes", handleIncomes)

	http.HandleFunc("/api/v1/incomes", handleAPIIncomes)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panicln()
	}
}
