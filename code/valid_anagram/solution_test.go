package valid_anagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s, t string
		want bool
	}{
		{
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			s:    "rat",
			t:    "cat",
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, IsAnagram(tt.s, tt.t))
	}
}
