package can_you_eat_your_favorite_candy_on_your_favorite_day

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanEat(t *testing.T) {
	tests := []struct {
		candies []int
		queries [][]int
		want    []bool
	}{
		{
			candies: []int{7, 4, 5, 3, 8},
			queries: [][]int{
				{0, 2, 2},
				{4, 2, 4},
				{2, 13, 1000000000},
			},
			want: []bool{true, false, true},
		},
		{
			candies: []int{5, 2, 6, 4, 1},
			queries: [][]int{
				{3, 1, 2},
				{4, 10, 3},
				{3, 10, 100},
				{4, 100, 30},
				{1, 3, 1},
			},
			want: []bool{false, true, true, false, false},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, CanEat(tt.candies, tt.queries))
	}
}
