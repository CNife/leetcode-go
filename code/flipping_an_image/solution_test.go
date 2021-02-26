package flipping_an_image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlipAndInvertImage(t *testing.T) {
	tests := []struct {
		image, want [][]int
	}{
		{
			image: [][]int{
				{1, 1, 0},
				{1, 0, 1},
				{0, 0, 0},
			},
			want: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
		},
		{
			image: [][]int{
				{1, 1, 0, 0},
				{1, 0, 0, 1},
				{0, 1, 1, 1},
				{1, 0, 1, 0},
			},
			want: [][]int{
				{1, 1, 0, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 1},
				{1, 0, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FlipAndInvertImage(tt.image))
	}
}
