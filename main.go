package main

import (
	"finnon/internal/database"
	"finnon/internal/domain"
	"log"
)

func main() {
	log.Println("Start Finnon")

	repo := database.NewIncomeRepositorySQLite()

	service := domain.NewIncomeService(repo)

	income, err := service.CreateIncome(1000, "teste002")
	if err != nil {
		log.Println("error with service and repository to save")
	}
	log.Println("now your repository save this incom:", income)

}
