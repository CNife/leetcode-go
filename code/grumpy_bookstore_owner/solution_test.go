package grumpy_bookstore_owner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSatisfied(t *testing.T) {
	tests := []struct {
		customers, grumpy []int
		x, want           int
	}{
		{
			customers: []int{1, 0, 1, 2, 1, 1, 7, 5},
			grumpy:    []int{0, 1, 0, 1, 0, 1, 0, 1},
			x:         3,
			want:      16,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MaxSatisfied(tt.customers, tt.grumpy, tt.x))
	}
}
