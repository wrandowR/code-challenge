package repository

//implementacion

import (
	"github.com/ansel1/merry"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/wrandowR/code-challenge/domain/model"
	"github.com/wrandowR/code-challenge/infrastructure/datastore"
	"github.com/wrandowR/code-challenge/usecase/repository"
)

var transactionsTable = goqu.T("transactions")

type transactionRepository struct {
	db                *goqu.Database
	transactionsTable exp.IdentifierExpression
}

var TransactionRepository repository.Transactions = &transactionRepository{
	db:                &datastore.SQLDBGoqu,
	transactionsTable: transactionsTable,
}

func (t *transactionRepository) Transactions(clientID string) ([]*model.Transaction, error) {

	return nil, merry.New("not implemented")
}

func (t *transactionRepository) SaveTransaction(transaction *model.Transaction) (*model.Transaction, error) {
	_, err := t.db.Insert(t.transactionsTable).Cols(
		"amount",
		"date",
	).Vals(goqu.Vals{
		transaction.Amount,
		transaction.Date,
	}).Executor().Exec()

	if err != nil {
		return nil, merry.Wrap(err)
	}

	return &model.Transaction{}, nil
}
