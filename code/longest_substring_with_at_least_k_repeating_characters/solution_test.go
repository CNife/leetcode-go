package longest_substring_with_at_least_k_repeating_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubstring(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s       string
		k, want int
	}{
		{
			s:    "aaabb",
			k:    3,
			want: 3,
		},
		{
			s:    "ababbc",
			k:    2,
			want: 5,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, LongestSubstring(tt.s, tt.k))
	}
}
