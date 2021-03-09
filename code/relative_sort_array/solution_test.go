package relative_sort_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelativeSortArray(t *testing.T) {
	tests := []struct {
		arr1, arr2, want []int
	}{
		{
			arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			arr2: []int{2, 1, 4, 3, 9, 6},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, RelativeSortArray(tt.arr1, tt.arr2))
	}
}
