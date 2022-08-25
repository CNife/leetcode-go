package find_k_closest_elements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindClosestElements(t *testing.T) {
	tests := []struct {
		arr  []int
		k, x int
		want []int
	}{
		{
			arr:  []int{1, 2, 3, 4, 5},
			k:    4,
			x:    3,
			want: []int{1, 2, 3, 4},
		},
		{
			arr:  []int{1, 2, 3, 4, 5},
			k:    4,
			x:    -1,
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FindClosestElements(tt.arr, tt.k, tt.x))
	}
}
