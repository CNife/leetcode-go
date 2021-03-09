package non_overlapping_intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      int
	}{
		{
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			want:      1,
		},
		{
			intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
			want:      2,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, EraseOverlapIntervals(tt.intervals))
	}
}
