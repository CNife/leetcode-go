package longest_repeating_character_replacement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterReplacement(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s       string
		k, want int
	}{
		{
			s:    "ABAB",
			k:    2,
			want: 4,
		},
		{
			s:    "AABABBA",
			k:    1,
			want: 4,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, CharacterReplacement(tt.s, tt.k))
	}
}
