package best_sightseeing_pair

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxScoreSightseeingPair(t *testing.T) {
	tests := []struct {
		array []int
		want  int
	}{
		{
			array: []int{8, 1, 5, 2, 6},
			want:  11,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MaxScoreSightseeingPair(tt.array))
	}
}
