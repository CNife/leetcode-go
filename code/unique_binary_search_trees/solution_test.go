package unique_binary_search_trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumTrees(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{
			n:    3,
			want: 5,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, NumTrees(tt.n))
	}
}
