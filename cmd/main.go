package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/ansel1/merry/v2"
)

//var fileTest := "test.csv"
/*
Total balance is 39.74
Number of transactions in July: 2
Number of transactions in August: 2

Average debit amount: -15.38
Average credit amount: 35.25
*/

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

	var totalBalance float64
	/*
		var NumbreOfTransactions int

	*/
	csvReader := csv.NewReader(r)

	var totalCreditTransactions float64
	var totalDebitTransactions float64

	header := false
	var AverageCreditAmountData []float64
	var AverageDebitAmountData []float64

	transactionsPerMonth := make(map[string]int)
	for {

		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if !header {
			header = true
			continue
		}

		cleanTransactionAmount, err := cleanAndParseTransaction(record[2])
		if err != nil {
			log.Fatal(err)
		}
		ok := isNegative(record[2])
		if ok {
			totalDebitTransactions += cleanTransactionAmount
			AverageDebitAmountData = append(AverageDebitAmountData, cleanTransactionAmount)
		} else {
			totalCreditTransactions += cleanTransactionAmount
			AverageCreditAmountData = append(AverageCreditAmountData, cleanTransactionAmount)
		}

		getTransactionsPerMonth(transactionsPerMonth, record[1])

	}

	creditAverage := average(AverageCreditAmountData)
	debitAverage := average(AverageDebitAmountData)

	totalBalance = math.Round((totalCreditTransactions-totalDebitTransactions)*100) / 100
	//return math.Round((a+b)*100) / 100
	//redondear esta vaina o hacerlo perfecto solo 2 digitos
	fmt.Println("total", totalBalance)

	fmt.Printf("average credit amount: %.2f\n", creditAverage)
	fmt.Printf("average debit amount: -%.2f\n", debitAverage)

	fmt.Println(transactionsPerMonth)
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

func getTransactionsPerMonth(monthMap map[string]int, date string) map[string]int {

	month := getMonth(date)

	if _, ok := monthMap[month]; ok {
		monthMap[month]++
		return monthMap
	}
	monthMap[month] = 1

	return monthMap
}
