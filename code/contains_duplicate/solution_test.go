package contains_duplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, ContainsDuplicate(tt.nums))
	}
}
