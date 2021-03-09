package add_to_array_form_of_integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToArrayForm(t *testing.T) {
	tests := []struct {
		array []int
		k     int
		want  []int
	}{
		{
			array: []int{1, 2, 0, 0},
			k:     34,
			want:  []int{1, 2, 3, 4},
		},
		{
			array: []int{2, 7, 4},
			k:     181,
			want:  []int{4, 5, 5},
		},
		{
			array: []int{2, 1, 5},
			k:     806,
			want:  []int{1, 0, 2, 1},
		},
		{
			array: []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			k:     1,
			want:  []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			array: []int{6},
			k:     809,
			want:  []int{8, 1, 5},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, AddToArrayForm(tt.array, tt.k))
	}
}
