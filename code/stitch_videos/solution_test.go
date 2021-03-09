package stitch_videos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVideoStitching(t *testing.T) {
	tests := []struct {
		clips [][]int
		t     int
		want  int
	}{
		{
			clips: [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}},
			t:     10,
			want:  3,
		},
		{
			clips: [][]int{{0, 1}, {1, 2}},
			t:     5,
			want:  -1,
		},
		{
			clips: [][]int{
				{0, 1},
				{6, 8},
				{0, 2},
				{5, 6},
				{0, 4},
				{0, 3},
				{6, 7},
				{1, 3},
				{4, 7},
				{1, 4},
				{2, 5},
				{2, 6},
				{3, 4},
				{4, 5},
				{5, 7},
				{6, 9},
			},
			t:    9,
			want: 3,
		},
		{
			clips: [][]int{{0, 4}, {2, 8}},
			t:     5,
			want:  2,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, VideoStitching(tt.clips, tt.t))
	}
}
