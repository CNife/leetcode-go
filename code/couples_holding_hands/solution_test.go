package couples_holding_hands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSwapsCouples(t *testing.T) {
	tests := []struct {
		row  []int
		want int
	}{
		{
			row:  []int{0, 2, 1, 3},
			want: 1,
		},
		{
			row:  []int{3, 2, 0, 1},
			want: 0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MinSwapsCouples(tt.row))
	}
}
