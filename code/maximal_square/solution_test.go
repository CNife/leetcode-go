package maximal_square

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximalSquare(t *testing.T) {
	tests := []struct {
		matrix [][]byte
		want   int
	}{
		{
			matrix: [][]byte{
				{'0', '0', '1', '0'},
				{'1', '1', '1', '1'},
				{'1', '1', '1', '1'},
				{'1', '1', '1', '0'},
				{'1', '1', '0', '0'},
				{'1', '1', '1', '1'},
				{'1', '1', '1', '0'},
			},
			want: 9,
		},
		{
			matrix: [][]byte{
				{'1', '0'},
				{'1', '0'},
			},
			want: 1,
		},
		{
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			want: 4,
		},
		{
			matrix: [][]byte{
				{'0', '1'},
				{'1', '0'},
			},
			want: 1,
		},
		{
			matrix: [][]byte{
				{'0'},
			},
			want: 0,
		},
		{
			matrix: [][]byte{
				{'1'},
			},
			want: 1,
		},
		{
			matrix: [][]byte{
				{'0', '0'},
				{'0', '0'},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MaximalSquare(tt.matrix))
	}
}
