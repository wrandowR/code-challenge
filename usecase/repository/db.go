package repository

import "github.com/wrandowR/code-challenge/domain/model"

type Transactions interface {
	Transactions(clientID string) ([]*model.Transaction, error)
	SaveTransaction(transaction *model.Transaction) (*model.Transaction, error)
}
