package word_pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordPattern(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		pattern, s string
		want       bool
	}{
		{
			pattern: "abc",
			s:       "dog cat dog",
			want:    false,
		},
		{
			pattern: "abba",
			s:       "dog cat cat dog",
			want:    true,
		},
		{
			pattern: "abba",
			s:       "dog cat cat fish",
			want:    false,
		},
		{
			pattern: "aaaa",
			s:       "dog cat cat dog",
			want:    false,
		},
		{
			pattern: "abba",
			s:       "dog dog dog dog",
			want:    false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, WordPattern(tt.pattern, tt.s))
	}
}
