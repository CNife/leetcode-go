package rabbits_in_forest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumRabbits(t *testing.T) {
	tests := []struct {
		answers []int
		want    int
	}{
		{
			answers: []int{1, 1, 2},
			want:    5,
		},
		{
			answers: []int{10, 10, 10},
			want:    11,
		},
		{
			answers: nil,
			want:    0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, NumRabbits(tt.answers))
	}
}
