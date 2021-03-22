package k_closest_points_to_origin

import (
	"testing"

	"github.com/CNife/leetcode-go/util/sort"
	"github.com/stretchr/testify/assert"
)

func TestKClosest(t *testing.T) {
	tests := []struct {
		points [][]int
		k      int
		want   [][]int
	}{
		{
			points: [][]int{{1, 3}, {-2, 2}},
			k:      1,
			want:   [][]int{{-2, 2}},
		},
		{
			points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
			k:      2,
			want:   [][]int{{3, 3}, {-2, 4}},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, sort.IntSlices(tt.want),
			sort.IntSlices(KClosest(tt.points, tt.k)))
	}
}
