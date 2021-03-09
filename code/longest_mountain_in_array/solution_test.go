package longest_mountain_in_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestMountain(t *testing.T) {
	tests := []struct {
		array []int
		want  int
	}{
		{
			array: []int{2, 1, 4, 7, 3, 2, 5},
			want:  5,
		},
		{
			array: []int{2, 2, 2},
			want:  0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, LongestMountain(tt.array))
	}
}
