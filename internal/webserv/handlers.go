package webserv

import (
	"encoding/json"
	"finnon/internal/database"
	"finnon/internal/domain"
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
		log.Println("hello GET")
		database.SelectIncomes()
	}
	if r.Method == http.MethodPost {
		log.Println("hello POST")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("error ao ler o body:", err)
		}

		var income *domain.Income
		err = json.Unmarshal(body, &income)
		if err != nil {
			log.Println("error to use unmarshall")
		}

		log.Println("Find:", income)

		database.InsertIncome(*income)
		log.Println("fim post")
	}
}