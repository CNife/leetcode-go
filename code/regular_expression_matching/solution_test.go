package regular_expression_matching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s, p string
		want bool
	}{
		{
			s:    "aa",
			p:    "a",
			want: false,
		},
		{
			s:    "aa",
			p:    "a*",
			want: true,
		},
		{
			s:    "ab",
			p:    ".*",
			want: true,
		},
		{
			s:    "aab",
			p:    "c*a*b",
			want: true,
		},
		{
			s:    "mississippi",
			p:    "mis*is*p*.",
			want: false,
		},
		{
			s:    "aa",
			p:    "*a",
			want: false,
		},
		{
			s:    "a",
			p:    "",
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, IsMatch(tt.s, tt.p))
	}
}
