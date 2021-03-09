package minimum_window_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s, t, want string
	}{
		{
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			s:    "a",
			t:    "a",
			want: "a",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MinWindow(tt.s, tt.t))
	}
}
