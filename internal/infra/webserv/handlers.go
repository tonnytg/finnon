package webserv

import (
	"encoding/json"
	"finnon/internal/domain"
	database2 "finnon/internal/infra/database"
	"fmt"
	"io"
	"log"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/favicon.ico" {
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodOptions {
		log.Println("hello OPTIONS")
		return
	}
	if r.Method == http.MethodGet {
		log.Println("hello GET")
		log.Println(r)
	}
	if r.Method == http.MethodPost {
		log.Println("hello POST")
		log.Println(r)
	}
}

func handleIncome(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		log.Println("hello OPTIONS")
		return
	}
	if r.Method == http.MethodGet {
		db := database2.NewDB(nil)
		incomes := database2.SelectIncomes(db)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		jsonIncomes, _ := json.Marshal(incomes)

		_, _ = w.Write(jsonIncomes)

	}
	if r.Method == http.MethodPost {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("error ao ler o body:", err)
		}

		var income *domain.Income
		err = json.Unmarshal(body, &income)
		if err != nil {
			log.Println("error to use unmarshall")
		}

		repo := database2.NewIncomeRepositorySQLite()
		service := domain.NewIncomeService(repo)

		_, err = service.CreateIncome(income.Description,
			income.Amount,
			income.Source,
			income.Provider,
			income.PaymentDate,
			income.Type)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprintf(w, "error")
			w.Header().Set("Content-Type", "application/json")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, "ok")
		return
	}
}
