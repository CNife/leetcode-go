package capacity_to_ship_packages_within_d_days

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShipWithinDays(t *testing.T) {
	tests := []struct {
		weights []int
		d, want int
	}{
		{
			weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			d:       5,
			want:    15,
		},
		{
			weights: []int{3, 2, 2, 4, 1, 4},
			d:       3,
			want:    6,
		},
		{
			weights: []int{1, 2, 3, 1, 1},
			d:       4,
			want:    3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, ShipWithinDays(tt.weights, tt.d))
	}
}
