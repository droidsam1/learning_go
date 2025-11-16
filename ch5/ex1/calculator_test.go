package calculator

import (
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
			got, err := perform(tt.input)
			assert.ErrorIs(t, err, tt.targetErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
