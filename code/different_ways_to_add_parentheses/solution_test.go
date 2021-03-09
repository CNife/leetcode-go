package different_ways_to_add_parentheses

import (
	"testing"

	"github.com/CNife/leetcode-go/util"
	"github.com/stretchr/testify/assert"
)

func TestDiffWaysToCompute(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{
			input: "2-1-1",
			want:  []int{0, 2},
		},
		{
			input: "2*3-4*5",
			want:  []int{-34, -14, -10, -10, 10},
		},
		{
			input: "11",
			want:  []int{11},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, util.SortInts(tt.want),
			util.SortInts(DiffWaysToCompute(tt.input)))
	}
}
