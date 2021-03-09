package get_equal_substrings_within_budget

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualSubstring(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s, t          string
		maxCost, want int
	}{
		{
			s:       "abcd",
			t:       "bcdf",
			maxCost: 3,
			want:    3,
		},
		{
			s:       "abcd",
			t:       "cdef",
			maxCost: 3,
			want:    1,
		},
		{
			s:       "abcd",
			t:       "acde",
			maxCost: 0,
			want:    1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, EqualSubstring(tt.s, tt.t, tt.maxCost))
	}
}
