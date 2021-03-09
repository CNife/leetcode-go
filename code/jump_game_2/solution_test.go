package jump_game_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 3, 1, 1, 4},
			want: 2,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Jump(tt.nums))
	}
}
