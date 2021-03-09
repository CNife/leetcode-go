package minimum_number_of_arrows_to_burst_balloons

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinArrowShots(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{
			points: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			want:   2,
		},
		{
			points: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			want:   4,
		},
		{
			points: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			want:   2,
		},
		{
			points: [][]int{{1, 2}},
			want:   1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FindMinArrowShots(tt.points))
	}
}
