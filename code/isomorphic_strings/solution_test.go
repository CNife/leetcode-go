package isomorphic_strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{
		{
			s:    "egg",
			t:    "add",
			want: true,
		},
		{
			s:    "foo",
			t:    "bar",
			want: false,
		},
		{
			s:    "paper",
			t:    "title",
			want: true,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, IsIsomorphic(tt.s, tt.t))
	}
}
