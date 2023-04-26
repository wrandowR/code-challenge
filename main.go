package main

import (
	"fmt"
	"log"

	"github.com/ansel1/merry"
	"github.com/wrandowR/code-challenge/config"
	"github.com/wrandowR/code-challenge/infrastructure/datastore"
	"github.com/wrandowR/code-challenge/usecase/interactor"
	"github.com/wrandowR/code-challenge/usecase/service"
)

func main() {
	config.ReadConfig()

	if err := datastore.NewDBConn(); err != nil {
		log.Fatal(merry.Wrap(err))
	}

	emailSender := service.NewEmailSender()

	processor := interactor.NewFileProcessor(nil, emailSender)

	if err := processor.ProccesFile("transactions.csv"); err != nil {
		log.Fatal(merry.Wrap(err))
	}

	fmt.Println("Done")
}
