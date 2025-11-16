package calculator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      int
		targetErr error
	}{
		{"sum two numbers", "1+1", 2, nil},
		{"sum two numbers with spaces", " 1 + 1 ", 2, nil},
		{"sum two numbers with spaces", "           1     +1 ", 2, nil},
		{"sum a negative number", "-1+1", 0, nil},
		{"substract two numbers", "2-1", 1, nil},
		{"substract a negative number", "-1-1", -2, nil},
		{"multiply a two numbers", "2*2", 4, nil},
		{"multiply two negative numbers", "-2*-2", 4, nil},
		{"divide two numbers", "2/2", 1, nil},
		{"divide by zero", "2/0", 0, ErrDivisionByZero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Perform(tt.input)
			assert.ErrorIs(t, err, tt.targetErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCalculatorInvalidInputs(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  error
	}{
		{"throw error when invalid input", "abc+1", errors.New("invalid left operand in \"abc+1\": strconv.Atoi: parsing \"abc\": invalid syntax")},
		{"throw error when invalid operand", "1/", errors.New("invalid right operand in \"1/\": strconv.Atoi: parsing \"\": invalid syntax")},
		{"throw error when no operator found", "", ErrUnsupportedOperator},
		{"throw error when operator only", "+", errors.New("invalid left operand in \"+\": strconv.Atoi: parsing \"\": invalid syntax")},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, got := Perform(tt.input)
			assert.EqualError(t, got, tt.want.Error())
		})
	}
}
