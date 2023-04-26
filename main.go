package main

import (
	"fmt"

	"github.com/wrandowR/code-challenge/domain/model"
	"github.com/wrandowR/code-challenge/usecase/service"
)

func main() {

	emailSender := service.NewEmailSender()

	transactionData := model.TransactionEmail{
		TotalBalance: 100,
		Transactions: []model.TransactionInAMounth{
			{
				Month: "January",
				Total: 100,
			},
			{
				Month: "February",
				Total: 100,
			},
		},
		AverageDebitAmount:  200,
		AverageCreditAmount: -10,
	}

	emailSender.SendEmail("test", &transactionData)
	/*
		processor := interactor.NewFileProcessor(nil, emailSender)

		if err := processor.ProccesFile("transactions.csv"); err != nil {
			log.Fatal(merry.Wrap(err))
		}*/

	fmt.Println("Done")
}
