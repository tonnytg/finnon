package webserver

import (
	"log"
	"net/http"
)

func Start() {

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/incomes", handleIncomes)
	http.HandleFunc("/outcomes", handleOutcomes)

	http.HandleFunc("/api/v1/incomes", handleAPIIncomes)
	http.HandleFunc("/income", handleIncome)
	http.HandleFunc("/api/v1/outcomes", handleAPIOutcomes)
	http.HandleFunc("/outcome", handleOutcome)

	log.Println("Start web server :8089")
	if err := http.ListenAndServe(":8089", nil); err != nil {
		log.Panicln()
	}
}
