package monotonic_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMonotonic(t *testing.T) {
	tests := []struct {
		array []int
		want  bool
	}{
		{
			array: []int{1, 2, 2, 3},
			want:  true,
		},
		{
			array: []int{6, 5, 4, 4},
			want:  true,
		},
		{
			array: []int{1, 3, 2},
			want:  false,
		},
		{
			array: []int{1, 2, 4, 5},
			want:  true,
		},
		{
			array: []int{1, 1, 1},
			want:  true,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, IsMonotonic(tt.array))
	}
}
