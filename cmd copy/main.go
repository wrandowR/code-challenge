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
		var AverageDebitAmount float64
		var AverageCreditAmount float64
	*/
	csvReader := csv.NewReader(r)

	var totalCreditTransactions float64
	var totalDebitTransactions float64

	header := false
	var AverageCreditAmountData []float64
	var AverageDebitAmountData []float64
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

	}

	creditAverage := average(AverageCreditAmountData)
	debitAverage := average(AverageDebitAmountData)

	totalBalance = math.Round((totalCreditTransactions-totalDebitTransactions)*100) / 100
	//return math.Round((a+b)*100) / 100
	//redondear esta vaina o hacerlo perfecto solo 2 digitos
	fmt.Println(totalBalance, creditAverage, debitAverage)
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

/*

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	var balance float64
	var credits float64
	var debits float64
	for scanner.Scan() {
		transaction := scanner.Text()
		amount, err := strconv.ParseFloat(strings.TrimSpace(transaction[1:]), 64)
		if err != nil {
			log.Printf("Error parsing transaction amount '%v': %v", transaction, err)
			continue
		}
		if transaction[0] == '+' {
			credits += amount
			balance += amount
		} else if transaction[0] == '-' {
			debits += amount
			balance -= amount
		} else {
			log.Printf("Invalid transaction '%v'", transaction)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Format the email body
	subject := "Transaction Summary"
	body := fmt.Sprintf("Credits: %.2f\nDebits: %.2f\nBalance: %.2f", credits, debits, balance)

	// Send the email
	err = sendEmail(subject, body)
	if err != nil {
		log.Fatalf("Error sending email: %v", err)
	}
	fmt.Println("Email sent successfully")
}*/
