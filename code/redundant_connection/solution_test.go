package redundant_connection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRedundantConnection(t *testing.T) {
	tests := []struct {
		edges [][]int
		want  []int
	}{
		{
			edges: [][]int{{1, 2}, {1, 3}, {2, 3}},
			want:  []int{2, 3},
		},
		{
			edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			want:  []int{1, 4},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FindRedundantConnection(tt.edges))
	}
}
