package smallest_string_with_swaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestStringWithSwaps(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s     string
		pairs [][]int
		want  string
	}{
		{
			s:     "dcab",
			pairs: [][]int{{0, 3}, {1, 2}},
			want:  "bacd",
		},
		{
			s:     "dcab",
			pairs: [][]int{{0, 3}, {1, 2}, {0, 2}},
			want:  "abcd",
		},
		{
			s:     "cba",
			pairs: [][]int{{0, 1}, {1, 2}},
			want:  "abc",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, SmallestStringWithSwaps(tt.s, tt.pairs))
	}
}
