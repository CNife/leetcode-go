package find_the_difference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheDifference(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s, t string
		want byte
	}{
		{
			s:    "abcd",
			t:    "abcde",
			want: 'e',
		},
		{
			s:    "",
			t:    "y",
			want: 'y',
		},
		{
			s:    "a",
			t:    "aa",
			want: 'a',
		},
		{
			s:    "ae",
			t:    "aea",
			want: 'a',
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FindTheDifference(tt.s, tt.t))
	}
}
