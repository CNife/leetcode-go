package first_unique_character_in_a_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstUnique(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "leetcode",
			want: 0,
		},
		{
			s:    "loveleetcode",
			want: 2,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FirstUnique(tt.s))
	}
}
