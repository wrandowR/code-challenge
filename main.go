package main

import (
	"fmt"
	"log"

	"github.com/ansel1/merry"
	"github.com/wrandowR/code-challenge/config"
	"github.com/wrandowR/code-challenge/usecase/interactor"
	"github.com/wrandowR/code-challenge/usecase/service"
)

func main() {
	config.ReadConfig()
	emailSender := service.NewEmailSender()

	processor := interactor.NewFileProcessor(nil, emailSender)

	if err := processor.ProccesFile("transactions.csv"); err != nil {
		log.Fatal(merry.Wrap(err))
	}

	fmt.Println("Done")
}
