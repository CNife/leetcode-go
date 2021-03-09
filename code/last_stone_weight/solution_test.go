package last_stone_weight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		stones []int
		want   int
	}{
		{
			stones: []int{2, 7, 4, 1, 8, 1},
			want:   1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, LastStoneWeight(tt.stones))
	}
}
