package number_of_operations_to_make_network_connected

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeConnected(t *testing.T) {
	tests := []struct {
		n           int
		connections [][]int
		want        int
	}{
		{
			n:           4,
			connections: [][]int{{0, 1}, {0, 2}, {1, 2}},
			want:        1,
		},
		{
			n:           6,
			connections: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}},
			want:        2,
		},
		{
			n:           6,
			connections: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}},
			want:        -1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MakeConnected(tt.n, tt.connections))
	}
}
