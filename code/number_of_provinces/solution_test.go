package number_of_provinces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCircleNum(t *testing.T) {
	tests := []struct {
		connects [][]int
		want     int
	}{
		{
			connects: nil,
			want:     0,
		},
		{
			connects: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			want: 2,
		},
		{
			connects: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FindCircleNum(tt.connects))
	}
}
