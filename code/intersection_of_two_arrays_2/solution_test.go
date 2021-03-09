package intersection_of_two_arrays_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		lhs, rhs, want []int
	}{
		{
			lhs:  []int{1, 2, 2, 1},
			rhs:  []int{2, 2},
			want: []int{2, 2},
		},
		{
			lhs:  []int{4, 9, 5},
			rhs:  []int{9, 4, 9, 8, 4},
			want: []int{4, 9},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, buildMap(tt.want), buildMap(Intersect(tt.lhs, tt.rhs)))
	}
}
