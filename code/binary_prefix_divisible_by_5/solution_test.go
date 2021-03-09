package binary_prefix_divisible_by_5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixDividedBy5(t *testing.T) {
	tests := []struct {
		array []int
		want  []bool
	}{
		{
			array: []int{0, 1, 1},
			want:  []bool{true, false, false},
		},
		{
			array: []int{1, 1, 1},
			want:  []bool{false, false, false},
		},
		{
			array: []int{0, 1, 1, 1, 1, 1},
			want:  []bool{true, false, false, false, true, false},
		},
		{
			array: []int{1, 1, 1, 0, 1},
			want:  []bool{false, false, false, false, false},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, PrefixDividedBy5(tt.array))
	}
}
