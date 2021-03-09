package triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		triangle [][]int
		want     int
	}{
		{
			triangle: [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MinimumTotal(tt.triangle))
	}
}
