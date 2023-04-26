package main

import (
	"fmt"

	"github.com/wrandowR/code-challenge/usecase/service"
)

func main() {

	emailSender := service.NewEmailSender()
	emailSender.SendEmail("test", "test")
	/*
		processor := interactor.NewFileProcessor(nil, emailSender)

		if err := processor.ProccesFile("transactions.csv"); err != nil {
			log.Fatal(merry.Wrap(err))
		}*/

	fmt.Println("Done")
}
