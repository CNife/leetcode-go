package matrix_cells_in_distance_order

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllCellsDistanceOrder(t *testing.T) {
	tests := []struct {
		r, c, r0, c0 int
		want         [][]int
	}{
		{
			r:  1,
			c:  2,
			r0: 0,
			c0: 0,
			want: [][]int{
				{0, 0},
				{0, 1},
			},
		},
		{
			r:  2,
			c:  2,
			r0: 0,
			c0: 1,
			want: [][]int{
				{0, 1},
				{0, 0},
				{1, 1},
				{1, 0},
			},
		},
		{
			r:  2,
			c:  3,
			r0: 1,
			c0: 2,
			want: [][]int{
				{1, 2},
				{0, 2},
				{1, 1},
				{0, 1},
				{1, 0},
				{0, 0},
			},
		},
	}
	for _, tt := range tests {
		assert.Equal(t,
			sortDistances(tt.want, tt.r0, tt.c0),
			sortDistances(
				AllCellsDistanceOrder(tt.r, tt.c, tt.r0, tt.c0),
				tt.r0,
				tt.c0,
			),
		)
	}
}

func sortDistances(points [][]int, x0, y0 int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		di := distance(points[i][0], points[i][1], x0, y0)
		dj := distance(points[j][0], points[j][1], x0, y0)
		if di == dj {
			if points[i][0] == points[j][0] {
				return points[i][1] < points[j][1]
			}
			return points[i][0] < points[j][0]
		}
		return di < dj
	})
	return points
}
