package largest_divisible_subset

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestDivisibleSubset(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 3},
			want: 2,
		},
		{
			nums: []int{1, 2, 4, 8},
			want: 4,
		},
	}
	for _, tt := range tests {
		got := LargestDivisibleSubset(tt.nums)
		assert.Equal(t, tt.want, len(got))

		sort.Ints(got)
		for i := 0; i < len(got)-1; i++ {
			for j := i + 1; j < len(got); j++ {
				assert.Equal(t, 0, got[j]%got[i])
			}
		}
	}
}
