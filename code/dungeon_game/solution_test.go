package dungeon_game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateMinimumHP(t *testing.T) {
	tests := []struct {
		dungeon [][]int
		want    int
	}{
		{
			dungeon: [][]int{
				{-2, -3, 3},
				{-5, -10, 1},
				{10, 30, -5},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, CalculateMinimumHP(tt.dungeon))
	}
}
