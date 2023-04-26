package main

import (
	"log"
	"os"

	"github.com/ansel1/merry/v2"
	"github.com/wrandowR/code-challenge/internal/service"
)

func main() {

	// Open the file
	file, err := os.Open("../test.csv")
	if err != nil {
		log.Fatal(merry.Append(err, "Error opening file"))
	}

	service := service.NewFileProcessorService()

	if err := service.ProccesFile(file); err != nil {
		log.Fatal(merry.Append(err, "Error processing file"))
	}

	defer file.Close()
}
