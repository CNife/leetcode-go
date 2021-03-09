package gas_station

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		gas, cost []int
		want      int
	}{
		{
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
			want: 3,
		},
		{
			gas:  []int{2, 3, 4},
			cost: []int{3, 4, 3},
			want: -1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, CanCompleteCircuit(tt.gas, tt.cost))
	}
}
