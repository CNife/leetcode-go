package n_queens_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalNQueens(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{
			n:    4,
			want: 2,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, TotalNQueens(tt.n))
	}
}
