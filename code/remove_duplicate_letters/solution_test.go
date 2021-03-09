package remove_duplicate_letters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s, want string
	}{
		{
			s:    "bcabc",
			want: "abc",
		},
		{
			s:    "cbacdcbc",
			want: "acdb",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, RemoveDuplicateLetters(tt.s))
	}
}
