package interactor

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTransactionsPerMonth(t *testing.T) {
	monthMap := make(map[string]int)

	// test for adding a new month
	result := getTransactionsPerMonth(monthMap, "January")
	assert.Equal(t, result["January"], 1)

	// test for adding another transaction to the same month
	result = getTransactionsPerMonth(monthMap, "January")
	assert.Equal(t, result["January"], 2)

	// test for adding a new month
	result = getTransactionsPerMonth(monthMap, "February")
	assert.Equal(t, result["February"], 1)
}

func TestIsNegative(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Negative input",
			input:    "-123",
			expected: true,
		},
		{
			name:     "Positive input",
			input:    "456",
			expected: false,
		},
		{
			name:     "Zero input",
			input:    "0",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isNegative(tt.input)
			if actual != tt.expected {
				t.Errorf("Unexpected result. input=%v, expected=%v, actual=%v", tt.input, tt.expected, actual)
			}
		})
	}
}

func TestCleanAndParseTransaction(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    float64
		expectedErr bool
		expectedMsg string
	}{
		{
			name:        "Valid transaction",
			input:       "12.34 USD",
			expected:    12.34,
			expectedErr: false,
			expectedMsg: "",
		},
		{
			name:        "Invalid transaction",
			input:       "abcd",
			expected:    0.0,
			expectedErr: true,
			expectedMsg: "strconv.ParseFloat: parsing \"\": invalid syntax",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := cleanAndParseTransaction(tt.input)
			if tt.expectedErr && err == nil {
				t.Errorf("Expected an error, but no error occurred")
			} else if !tt.expectedErr && err != nil {
				t.Errorf("Unexpected error occurred: %v", err)
			} else if tt.expectedErr && err != nil && err.Error() != tt.expectedMsg {
				t.Errorf("Unexpected error message. expected=%v, actual=%v", tt.expectedMsg, err.Error())
			} else if actual != tt.expected {
				t.Errorf("Unexpected result. input=%v, expected=%v, actual=%v", tt.input, tt.expected, actual)
			}
		})
	}
}

func TestRamdonDataTest(t *testing.T) {
	rand.Seed(time.Now().UnixNano()) // Semilla aleatoria para generar valores aleatorios

	// Bucle para generar 4 filas de datos de prueba
	for i := 0; i < 10000; i++ {
		// Generar un ID de transacción aleatorio entre 0 y 99
		transactionID := i

		// Generar una fecha aleatoria entre el 1 de julio y el 31 de agosto
		startDate := time.Date(2022, time.July, 1, 0, 0, 0, 0, time.UTC).Unix()
		endDate := time.Date(2022, time.August, 31, 23, 59, 59, 0, time.UTC).Unix()
		seconds := rand.Int63n(endDate-startDate) + startDate
		date := time.Unix(seconds, 0).Format("1/2")

		// Generar un valor aleatorio entre -100 y 100 con una precisión de dos decimales
		amount := rand.Float64()*200 - 100
		amountStr := strconv.FormatFloat(amount, 'f', 2, 64)

		// Generar una cadena que indique si el valor es positivo o negativo
		var sign string
		if amount >= 0 {
			sign = "+"
		}

		// Imprimir los datos de prueba generados
		fmt.Printf("%v,%v,%v%v\n", transactionID, date, sign, amountStr)
		t.Fail()
	}
}
