package minimum_cost_to_move_chips_to_the_same_position

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostToMoveChips(t *testing.T) {
	tests := []struct {
		position []int
		want     int
	}{
		{
			position: []int{1, 2, 3},
			want:     1,
		},
		{
			position: []int{2, 2, 2, 3, 3},
			want:     2,
		},
		{
			position: []int{1, 1000000000},
			want:     1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MinCostToMoveChips(tt.position))
	}
}
