package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			result, error := tt.input.toJSON()
			require.NoError(t, error)
			assert.Equal(t, tt.expected, result)
		})
	}

}
