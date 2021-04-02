package volume_of_histogram_lcci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Trap(tt.height))
	}
}
