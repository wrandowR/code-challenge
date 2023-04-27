package repository

//implementacion

import (
	"context"
	"time"

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

func (t *transactionRepository) GetTransactions(clientID string) ([]*model.Transaction, error) {

	return nil, merry.New("not implemented")
}

func (t *transactionRepository) SaveTransactions(transactions []*model.Transaction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return merry.Wrap(err)
	}

	query := goqu.Insert(t.transactionsTable).Rows(transactions)
	sql, args, err := query.ToSQL()
	if err != nil {
		return merry.Wrap(err)
	}

	_, err = tx.ExecContext(ctx, sql, args...)
	if err != nil {
		tx.Rollback()
		return merry.Wrap(err)

	}

	//comit tx with rollback
	if err = tx.Commit(); err != nil {
		return merry.Wrap(err)
	}

	return nil
}
