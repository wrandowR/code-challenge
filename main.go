package main

import (
	"fmt"

	"github.com/ansel1/merry"
	"github.com/sirupsen/logrus"
	"github.com/wrandowR/code-challenge/config"
	"github.com/wrandowR/code-challenge/infrastructure/datastore"
	repo "github.com/wrandowR/code-challenge/interface/repository"
	"github.com/wrandowR/code-challenge/usecase/interactor"
	"github.com/wrandowR/code-challenge/usecase/service"
)

var destinationEmail = "testemailchallenge@hotmail.com"
var fileName = "transactions.csv"

func main() {
	if err := config.ReadConfig(); err != nil {
		fmt.Println("Error reading config", err)
		logrus.Fatal(merry.Wrap(err))
	}

	if err := datastore.NewDBConn(); err != nil {
		fmt.Println("Error connecting to database", err)
		logrus.Fatal(merry.Wrap(err))
	}

	if config.EnableMigrations() {
		datastore.ResetDatabase()
		datastore.DoMigration()
	}

	emailSender, err := service.NewEmailSender(destinationEmail)
	if err != nil {
		logrus.Fatal(merry.Wrap(err))
	}

	processor := interactor.NewFileProcessor(repo.TransactionRepository, emailSender)

	if err := processor.ProccesFile(fileName); err != nil {
		logrus.Fatal(merry.Wrap(err))
	}

	fmt.Println("Done")
}
