package models

import (
	"testing"

	"yosbomb.com/bucketbudget/internal/assert"
)

func TestAmountAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        Amount
		b        Amount
		expected Amount
	}{
		{
			name:     "$5 + $5 = $10",
			a:        Amount{Cents: 500},
			b:        Amount{Cents: 500},
			expected: Amount{Cents: 1000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.a.Add(tt.b)

			assert.Equal(t, result, tt.expected)
		})
	}
}
