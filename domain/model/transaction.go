package model

type Transaction struct {
	ID     string
	Amount float64
	Date   string
}

// structura para los datos de las transacciones que se enviaran por email
type TransactionEmail struct {
	TotalBalance        float64
	Transactions        []TransactionInAMounth
	AverageDebitAmount  float64
	AverageCreditAmount float64
}

type TransactionInAMounth struct {
	Month string
	Total float64
}
