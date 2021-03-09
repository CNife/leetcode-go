package increasing_decreasing_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortString(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s, want string
	}{
		{
			s:    "aaaabbbbcccc",
			want: "abccbaabccba",
		},
		{
			s:    "rat",
			want: "art",
		},
		{
			s:    "leetcode",
			want: "cdelotee",
		},
		{
			s:    "ggggggg",
			want: "ggggggg",
		},
		{
			s:    "spo",
			want: "ops",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, SortString(tt.s))
	}
}
