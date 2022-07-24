package distance_between_bus_stops

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistanceBetweenBusStops(t *testing.T) {
	tests := []struct {
		distance                 []int
		start, destination, want int
	}{
		{
			distance:    []int{1, 2, 3, 4},
			start:       0,
			destination: 1,
			want:        1,
		},
		{
			distance:    []int{1, 2, 3, 4},
			start:       0,
			destination: 2,
			want:        3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, DistanceBetweenBusStops(tt.distance, tt.start, tt.destination))
	}
}
