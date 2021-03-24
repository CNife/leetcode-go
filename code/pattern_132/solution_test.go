package pattern_132

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind132Pattern(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			nums: []int{3, 1, 2, 4},
			want: true,
		},
		{
			nums: []int{-1, 3, 2, 0},
			want: true,
		},
		{
			nums: []int{1, 0, 1, -4, -3},
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Find132Pattern(tt.nums))
	}
}
