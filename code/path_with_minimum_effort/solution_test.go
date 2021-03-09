package path_with_minimum_effort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumEffortPath(t *testing.T) {
	tests := []struct {
		heights [][]int
		want    int
	}{
		{
			heights: [][]int{
				{1, 2, 2},
				{3, 8, 2},
				{5, 3, 5},
			},
			want: 2,
		},
		{
			heights: [][]int{
				{1, 2, 3},
				{3, 8, 4},
				{5, 3, 5},
			},
			want: 1,
		},
		{
			heights: [][]int{
				{1, 2, 1, 1, 1},
				{1, 2, 1, 2, 1},
				{1, 2, 1, 2, 1},
				{1, 2, 1, 2, 1},
				{1, 1, 1, 2, 1},
			},
			want: 0,
		},
		{
			heights: [][]int{{3}},
			want:    0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MinimumEffortPath(tt.heights))
	}
}
