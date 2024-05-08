package webserv

import (
	"log"
	"net/http"
)

func Start() {

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/incomes", handleIncomes)
	http.HandleFunc("/outcomes", handleOutcomes)

	http.HandleFunc("/api/v1/incomes", handleAPIIncomes)
	http.HandleFunc("/api/v1/outcomes", handleAPIOutcomes)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panicln()
	}
}
