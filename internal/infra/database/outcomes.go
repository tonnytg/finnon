package database

import (
	"finnon/internal/domain"
	"log"
)

func SelectOutcomes(db *DB) []domain.Outcome {

	conn := db

	var outcomes []domain.Outcome

	rows, err := conn.Db.Query("SELECT * FROM outcomes")
	if err != nil {
		log.Println("error to get information from clients")
		return nil
	}
	defer rows.Close()

	for rows.Next() {

		var outcome domain.Outcome

		err = rows.Scan(
			&outcome.ID,
			&outcome.Description,
			&outcome.Amount,
			&outcome.Tax,
			&outcome.IsPayed,
			&outcome.CompanyName,
			&outcome.Installment,
			&outcome.TotalInstallment,
			&outcome.PaidAt,
			&outcome.Category,
			&outcome.TypeOutcome,
			&outcome.Repeat,
			&outcome.CreatedAt)
		if err != nil {
			log.Println("panic:", err)
		}
		log.Println("Outcome:", outcome)
		outcomes = append(outcomes, outcome)
	}
	if err = rows.Err(); err != nil {
		log.Println("error to read rows:", err)
	}
	return outcomes
}
