package repository

import "github.com/wrandowR/code-challenge/domain/model"

type Transactions interface {
	GetTransactions(clientID string) ([]*model.Transaction, error)
	SaveTransactions(transaction []*model.Transaction) error
}
