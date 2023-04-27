package interactor

import (
	"testing"

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
