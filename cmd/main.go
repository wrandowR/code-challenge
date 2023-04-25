package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/ansel1/merry/v2"
)

var MAX_GOROUTINES = 10

type aux struct {
	isNegative bool
	amount     float64
	month      string
}

func main() {

	// Open the file
	file, err := os.Open("../test.csv")
	if err != nil {
		log.Fatal(merry.Append(err, "Error opening file"))
	}

	ReadFile(file)
	defer file.Close()

}
func ReadFile(r io.Reader) {

	csvReader := csv.NewReader(r)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//elimina el header
	records = records[1:]

	var totalBalance float64
	var totalCreditTransactions float64
	var totalDebitTransactions float64
	var AverageCreditAmountData []float64
	var AverageDebitAmountData []float64

	// Create worker pool
	var wg sync.WaitGroup
	jobs := make(chan []string, len(records))
	results := make(chan aux, len(records))

	for i := 0; i < MAX_GOROUTINES; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	// Send jobs to workers
	for _, record := range records {
		jobs <- record
	}
	close(jobs)

	monthMap := make(map[string]int)

	// Collect results from workers
	for i := 0; i < len(records); i++ {
		result := <-results

		getTransactionsPerMonth(monthMap, result.month)

		if result.isNegative {
			totalBalance -= result.amount
			totalDebitTransactions += result.amount
			AverageDebitAmountData = append(AverageDebitAmountData, result.amount)
			continue
		}

		totalBalance += result.amount
		totalCreditTransactions += result.amount
		AverageCreditAmountData = append(AverageCreditAmountData, result.amount)

	}

	wg.Wait()

	fmt.Println("total Balance", totalBalance)
	fmt.Println("total Debit Transactions", totalDebitTransactions)
	fmt.Println("total Credit Transactions", totalCreditTransactions)
	fmt.Println("Average Debit Amount Data", average(AverageDebitAmountData))
	fmt.Println("Average Credit Amount Data", average(AverageCreditAmountData))
	fmt.Println("Number of transactions month", monthMap)

}

func worker(jobs <-chan []string, results chan<- aux, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {

		cleanTransactionAmount, err := cleanAndParseTransaction(job[2])
		if err != nil {
			log.Fatal(err)
		}
		ok := isNegative(job[2])

		//guardar en base de datos

		results <- aux{
			isNegative: ok,
			amount:     cleanTransactionAmount,
			month:      getMonth(job[1]),
		}
	}
}

// funcion que tetorna si un string de numer so es negativo o positivoas
func isNegative(s string) bool {
	return s[0] == '-'
}

func average(numbers []float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func cleanAndParseTransaction(transaction string) (float64, error) {
	//remove the non-numeric characters from the string
	transaction = strings.Map(func(r rune) rune {
		if (r >= '0' && r <= '9') || r == '.' {
			return r
		}
		return -1
	}, transaction)

	//convert the string to a decimal number
	value, err := strconv.ParseFloat(transaction, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func getMonth(date string) string {
	parts := strings.Split(date, "/")
	month, _ := strconv.Atoi(parts[0])
	months := []string{
		"Enero",
		"Febrero",
		"Marzo",
		"Abril",
		"Mayo",
		"Junio",
		"Julio",
		"Agosto",
		"Septiembre",
		"Octubre",
		"Noviembre",
		"Diciembre"}

	return months[month-1]
}

func getTransactionsPerMonth(monthMap map[string]int, month string) map[string]int {

	if _, ok := monthMap[month]; ok {
		monthMap[month]++
		return monthMap
	}
	monthMap[month] = 1

	return monthMap
}
