package non_decreasing_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPossibility(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{3, 4, 2, 3},
			want: false,
		},
		{
			nums: []int{4, 2, 3},
			want: true,
		},
		{
			nums: []int{4, 2, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, CheckPossibility(tt.nums))
	}
}
