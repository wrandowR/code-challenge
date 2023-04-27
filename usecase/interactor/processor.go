package interactor

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ansel1/merry/v2"
	"github.com/sirupsen/logrus"
	"github.com/wrandowR/code-challenge/domain/model"
	repository "github.com/wrandowR/code-challenge/usecase/repository"
	"github.com/wrandowR/code-challenge/usecase/service"
)

// fileProcessor is a service that process a csv file
type fileProcessor struct {
	DataStore   repository.Transactions
	EmailSender service.EmailSender
}

func NewFileProcessor(dataStorage repository.Transactions, emailSender service.EmailSender) *fileProcessor {
	return &fileProcessor{
		DataStore:   dataStorage,
		EmailSender: emailSender,
	}
}

func (s *fileProcessor) ProccesFile(filename string) error {

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return merry.Wrap(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	monthMap := make(map[string]int)

	var totalCreditTransactionsAmount float64
	var totalDebitTransactionsAmount float64
	var totalCreditTransactions float64
	var totalDebitTransactions float64

	var totalRecords int
	header := true
	values := make([]*model.Transaction, 0)

	fmt.Println("started", time.Now())
	start := time.Now()
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if len(record) == 0 {
			break
		}
		if header {
			header = false
			continue
		}

		totalRecords++

		getTransactionsPerMonth(monthMap, getMonth(record[1]))

		cleanTransactionAmount, err := cleanAndParseTransaction(record[2])
		if err != nil {
			logrus.WithError(err).Error("error parsing transaction amount", record[2])
			continue
		}
		ok := isNegative(record[2])
		if ok {
			cleanTransactionAmount = cleanTransactionAmount * -1
			totalDebitTransactionsAmount += cleanTransactionAmount
			totalDebitTransactions++

		} else {
			totalCreditTransactionsAmount += cleanTransactionAmount
			totalCreditTransactions++
		}

		values = append(values, &model.Transaction{
			ID:     record[0],
			Amount: cleanTransactionAmount,
			Date:   record[1],
		})

		if len(values) == 1000 {
			//guardar en base de datos aca deberia estar guardando los datos asosiado a un id de usuario o cuenta, pero como no tengo datos,
			//prefiero dejarlo abierto , solo se muentras el ejemplo de un guardado en base de datos
			err := s.DataStore.SaveTransactions(values)
			if err != nil {
				logrus.WithError(err)
			}
			values = values[:0]
		}

	}

	if len(values) > 0 {
		err := s.DataStore.SaveTransactions(values)
		if err != nil {
			logrus.WithError(err)
		}
	}

	totalBalance := totalCreditTransactionsAmount + totalDebitTransactionsAmount
	avergeCreditAmount := totalCreditTransactionsAmount / totalCreditTransactions
	avergeDebitAmount := totalDebitTransactionsAmount / totalDebitTransactions

	TransactionInAMounth := []model.TransactionInAMounth{}
	for key, value := range monthMap {
		transactions := model.TransactionInAMounth{
			Month: key,
			Total: float64(value),
		}
		TransactionInAMounth = append(TransactionInAMounth, transactions)
	}

	transactionEmailData := model.TransactionEmail{
		TotalBalance:        math.Round(totalBalance*100) / 100,
		Transactions:        TransactionInAMounth,
		AverageDebitAmount:  math.Round(avergeDebitAmount*100) / 100,
		AverageCreditAmount: math.Round(avergeCreditAmount*100) / 100,
	}

	if err := s.EmailSender.SendEmail(&transactionEmailData); err != nil {
		return merry.Wrap(err)
	}

	elapsed := time.Since(start)
	fmt.Printf("Se procesaron %d registros en %v segundos\n", totalRecords, elapsed.Seconds())
	logrus.Info("Email sent")
	return nil
}

func isNegative(s string) bool {
	return s[0] == '-'
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
