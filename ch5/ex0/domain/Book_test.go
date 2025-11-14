package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJson(t *testing.T) {
	tests := []struct {
		testName string
		input    Book
		expected string
	}{
		{"Book marshalling to json", Book{Title: "Meditations", Author: "Marcus Aurelius", Pages: 254}, `{"title":"Meditations","author":"Marcus Aurelius","pages":254}`},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.ToJson())
		})
	}

}
