package service

import (
	"testing"

	"github.com/wrandowR/code-challenge/domain/model"
)

func TestParseTemplate(t *testing.T) {
	// Mock data
	data := &model.TransactionEmail{
		TotalBalance: 100.0,
		Transactions: []model.TransactionInAMounth{
			{
				Month: "January",
				Total: 100.0,
			},
		},
		AverageDebitAmount:  100.0,
		AverageCreditAmount: 100.0,
	}

	// Mock email sender
	emailSender := &emailSender{
		TransactionEmailTemplate: "email.html",
	}

	// Call function
	result, err := emailSender.parseTemplate(data)

	// Check for errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if result is not empty
	if result == "" {
		t.Errorf("Result is empty")
	}
}
