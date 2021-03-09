package valid_mountain_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidMountainArray(t *testing.T) {
	tests := []struct {
		array []int
		want  bool
	}{
		{
			array: []int{2, 1},
			want:  false,
		},
		{
			array: []int{3, 5, 5},
			want:  false,
		},
		{
			array: []int{0, 3, 2, 1},
			want:  true,
		},
		{
			array: []int{0, 1, 2, 3, 2, 1, 2, 1},
			want:  false,
		},
		{
			array: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want:  false,
		},
		{
			array: []int{0, 1, 2, 1, 2},
			want:  false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, ValidMountainArray(tt.array))
	}
}
