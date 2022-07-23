package offer_rebuild_sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequenceReconstruction(t *testing.T) {
	tests := []struct {
		nums      []int
		sequences [][]int
		want      bool
	}{
		{
			nums:      []int{1, 2, 3},
			sequences: [][]int{{1, 2}, {1, 3}},
			want:      false,
		},
		{
			nums:      []int{1, 2, 3},
			sequences: [][]int{{1, 2}},
			want:      false,
		},
		{
			nums:      []int{1, 2, 3},
			sequences: [][]int{{1, 2}, {1, 3}, {2, 3}},
			want:      true,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, SequenceReconstruction(tt.nums, tt.sequences))
	}
}
