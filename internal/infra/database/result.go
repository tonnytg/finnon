package database

type Result struct {
	Incomes  float64 `json:"incomes"`
	Outcomes float64 `json:"outcomes"`
	Result   float64 `json:"result"`
}

func SelectResult(db *DB) Result {

	query := `SELECT
    (SELECT COALESCE(SUM(amount), 0) FROM incomes) AS total_income,
    (SELECT COALESCE(SUM(amount), 0) FROM outcomes) AS total_outcome,
    (SELECT COALESCE(SUM(amount), 0) FROM incomes) - (SELECT COALESCE(SUM(amount), 0) FROM outcomes) AS balance;`

	rows, err := db.Db.Query(query)
	if err != nil {
		return Result{}
	}

	var incomes float64
	var outcomes float64
	var balance float64

	for rows.Next() {
		err = rows.Scan(&incomes, &outcomes, &balance)
		if err != nil {
			return Result{}
		}
	}

	result := Result{
		Incomes:  incomes,
		Outcomes: outcomes,
		Result:   balance,
	}
	return result
}
