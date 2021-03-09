package min_cost_climbing_stairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		costs []int
		want  int
	}{
		{
			costs: []int{10, 15, 20},
			want:  15,
		},
		{
			costs: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want:  6,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MinCostClimbingStairs(tt.costs))
	}
}
