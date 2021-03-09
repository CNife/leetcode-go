package target_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTargetSumWays(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{1, 1, 1, 1, 1},
			target: 3,
			want:   5,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FindTargetSumWays(tt.nums, tt.target))
	}
}
