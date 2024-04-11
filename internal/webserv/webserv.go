package webserv

import (
	"log"
	"net/http"
)

func Start() {

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/incomes", handleIncome)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panicln()
	}
}
