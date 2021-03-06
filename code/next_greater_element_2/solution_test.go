package next_greater_element_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreaterElements(t *testing.T) {
	tests := []struct {
		nums, want []int
	}{
		{
			nums: []int{1, 2, 1},
			want: []int{2, -1, 2},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, NextGreaterElements(tt.nums))
	}
}
