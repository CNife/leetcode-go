package remove_max_number_of_edges_to_keep_graph_fully_traversable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxNumEdgesToRemove(t *testing.T) {
	tests := []struct {
		n     int
		edges [][]int
		want  int
	}{
		{
			n: 4,
			edges: [][]int{
				{3, 1, 2},
				{3, 2, 3},
				{1, 1, 3},
				{1, 2, 4},
				{1, 1, 2},
				{2, 3, 4},
			},
			want: 2,
		},
		{
			n: 4,
			edges: [][]int{
				{3, 1, 2},
				{3, 2, 3},
				{1, 1, 4},
				{2, 1, 4},
			},
			want: 0,
		},
		{
			n: 4,
			edges: [][]int{
				{3, 2, 3},
				{1, 1, 2},
				{2, 3, 4},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MaxNumEdgesToRemove(tt.n, tt.edges))
	}
}
