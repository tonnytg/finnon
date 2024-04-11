package domain

type IncomeRepositoryInterface interface {
	Save(*Income) error
}
