package remove_duplicates_from_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums     []int
		want     int
		wantNums []int
	}{
		{
			nums:     []int{1, 1, 2},
			want:     2,
			wantNums: []int{1, 2},
		},
		{
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want:     5,
			wantNums: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, RemoveDuplicates(tt.nums))
		assert.Equal(t, tt.wantNums, tt.nums[:tt.want])
	}
}
