package beautiful_arrangement_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructArray(t *testing.T) {
	tests := []struct {
		n, k int
		want []int
	}{
		{
			n:    3,
			k:    1,
			want: []int{1, 2, 3},
		},
		{
			n:    3,
			k:    2,
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, ConstructArray(tt.n, tt.k))
	}
}
