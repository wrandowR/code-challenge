package service

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"sync"

	"code-challenge/config"
	"code-challenge/domain/model"

	"github.com/ansel1/merry/v2"
)

// FileProcessorService is a service that process a csv file
type fileProcessorService struct {
	//por mejorar :3
	DataStore   string
	EmailSender string
}

func NewFileProcessorService() *fileProcessorService {
	return &fileProcessorService{}
}

func (s *fileProcessorService) ProccesFile(r io.Reader) error {

	csvReader := csv.NewReader(r)

	records, err := csvReader.ReadAll()
	if err != nil {
		return merry.Wrap(err)
	}

	//delete header
	records = records[1:]

	var totalBalance float64
	var totalCreditTransactions float64
	var totalDebitTransactions float64
	var AverageCreditAmountData []float64
	var AverageDebitAmountData []float64

	// Create worker pool
	var wg sync.WaitGroup
	jobs := make(chan []string, len(records))
	results := make(chan model.Transaction, len(records))

	for i := 0; i < config.MaxGoroutines(); i++ {
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

		getTransactionsPerMonth(monthMap, result.Date)

		if result.IsNegative {
			totalBalance -= result.Amount
			totalDebitTransactions += result.Amount
			AverageDebitAmountData = append(AverageDebitAmountData, result.Amount)
			continue
		}

		totalBalance += result.Amount
		totalCreditTransactions += result.Amount
		AverageCreditAmountData = append(AverageCreditAmountData, result.Amount)

	}

	wg.Wait()

	//enviar correo aca con resultados

	fmt.Println("total Balance", totalBalance)
	fmt.Println("total Debit Transactions", totalDebitTransactions)
	fmt.Println("total Credit Transactions", totalCreditTransactions)
	fmt.Println("Average Debit Amount Data", average(AverageDebitAmountData))
	fmt.Println("Average Credit Amount Data", average(AverageCreditAmountData))
	fmt.Println("Number of transactions month", monthMap)

}

func worker(jobs <-chan []string, results chan<- model.Transaction, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {

		cleanTransactionAmount, err := cleanAndParseTransaction(job[2])
		if err != nil {
			log.Fatal(err)
		}
		ok := isNegative(job[2])

		//guardar en base de datos

		results <- model.Transaction{
			IsNegative: ok,
			Amount:     cleanTransactionAmount,
			Date:       getMonth(job[1]),
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
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December"}

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
