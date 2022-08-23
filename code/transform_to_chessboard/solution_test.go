package transform_to_chessboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovesToChessboard(t *testing.T) {
	tests := []struct {
		board [][]int
		want  int
	}{
		{
			board: [][]int{
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{1, 0, 0, 1},
				{1, 0, 0, 1},
			},
			want: 2,
		},
		{
			board: [][]int{
				{0, 1},
				{1, 0},
			},
			want: 0,
		},
		{
			board: [][]int{
				{1, 0},
				{1, 0},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MovesToChessboard(tt.board))
	}
}
