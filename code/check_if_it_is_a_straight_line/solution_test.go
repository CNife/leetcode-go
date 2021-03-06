package check_if_it_is_a_straight_line

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckStraightLine(t *testing.T) {
	tests := []struct {
		coordinates [][]int
		want        bool
	}{
		{
			coordinates: [][]int{{1, 2}, {2, 3}, {3, 5}},
			want:        false,
		},
		{
			coordinates: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}},
			want:        true,
		},
		{
			coordinates: [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}},
			want:        false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, CheckStraightLine(tt.coordinates))
	}
}
