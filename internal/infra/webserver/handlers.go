package webserver

import (
	"encoding/json"
	"finnon/internal/domain"
	"finnon/internal/infra/database"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.New("index.html").ParseFiles(
		"./internal/infra/webserver/templates/index.html"))

	resultData := database.SelectResult(database.NewDB(nil))

	if err := tmpl.Execute(w, resultData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if r.URL.Path == "/favicon.ico" {
		// TODO: what do you want to do with favicon?
		return
	}

	if r.Method == http.MethodOptions {
		log.Println("method OPTIONS")
		// TODO: what do you want to do with OPTIONS
		return
	}
	if r.Method == http.MethodGet {
		log.Println("method GET")
		// TODO: what do you want to do with GET
		log.Println(r)
	}
	if r.Method == http.MethodPost {
		log.Println("method POST")
		// TODO: what do you want to do with POST
		log.Println(r)
	}
}

func handleIncomes(w http.ResponseWriter, r *http.Request) {

	db := database.NewDB(nil)
	incomes := database.SelectIncomes(db)

	tmpl := template.Must(template.New("incomes.html").ParseFiles(
		"./internal/infra/webserver/templates/incomes.html"))
	if err := tmpl.Execute(w, incomes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if r.URL.Path == "/favicon.ico" {
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodOptions {
		log.Println("method OPTIONS")
		return
	}
	if r.Method == http.MethodGet {
		log.Println("method GET")
	}
	if r.Method == http.MethodPost {
		log.Println("method POST")
	}
}

func handleOutcomes(w http.ResponseWriter, r *http.Request) {

	db := database.NewDB(nil)
	outcomes := database.SelectOutcomes(db)

	tmpl := template.Must(template.New("outcomes.html").ParseFiles(
		"./internal/infra/webserver/templates/outcomes.html"))
	if err := tmpl.Execute(w, outcomes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if r.URL.Path == "/favicon.ico" {
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodOptions {
		log.Println("method OPTIONS")
		return
	}
	if r.Method == http.MethodGet {
		log.Println("method GET")
	}
	if r.Method == http.MethodPost {
		log.Println("method POST")
	}
}

func handleIncome(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.New("income_form.html").ParseFiles(
		"./internal/infra/webserver/templates/income_form.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func handleAPIIncomes(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		log.Println("hello OPTIONS")
		return
	}
	if r.Method == http.MethodGet {
		db := database.NewDB(nil)
		incomes := database.SelectIncomes(db)
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

		repo := database.NewIncomeRepositorySQLite()
		service := domain.NewIncomeService(repo)

		_, err = service.CreateIncome(income.Description,
			income.Amount,
			income.Source,
			income.Provider,
			income.PaymentDate,
			income.TypeIncome)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprintf(w, "error")
			w.Header().Set("Content-Type", "application/json")
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, "{\"message\":\"created\"}")
		return
	}
}

func handleOutcome(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.New("outcome_form.html").ParseFiles(
		"./internal/infra/webserver/templates/outcome_form.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func handleAPIOutcomes(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		log.Println("hello OPTIONS")
		return
	}
	if r.Method == http.MethodGet {
		db := database.NewDB(nil)
		outcomes := database.SelectOutcomes(db)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		jsonOutcomes, _ := json.Marshal(outcomes)

		_, _ = w.Write(jsonOutcomes)

	}
	if r.Method == http.MethodPost {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("error ao ler o body:", err)
		}

		var outcome *domain.Outcome
		err = json.Unmarshal(body, &outcome)
		if err != nil {
			log.Println("error to use unmarshall")
		}

		repo := database.NewOutcomeRepositorySQLite()
		service := domain.NewOutcomeService(repo)

		_, err = service.CreateOutcome(outcome.Description,
			outcome.Amount,
			outcome.Tax,
			outcome.IsPayed,
			outcome.CompanyName,
			outcome.Installment,
			outcome.TotalInstallment,
			outcome.PaidAt,
			outcome.Category,
			outcome.TypeOutcome,
			outcome.Repeat)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprintf(w, "{\"message\":\"error to create outcome\"}")
			w.Header().Set("Content-Type", "application/json")
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, "{\"message\":\"created\"}")
		return
	}
}
