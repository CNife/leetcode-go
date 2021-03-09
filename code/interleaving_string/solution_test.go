package interleaving_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInterleave(t *testing.T) {
	//noinspection SpellCheckingInspection
	tests := []struct {
		lhs, rhs, target string
		want             bool
	}{
		{
			lhs:    "aabcc",
			rhs:    "dbbca",
			target: "aadbbcbcac",
			want:   true,
		},
		{
			lhs:    "aabcc",
			rhs:    "dbbca",
			target: "aadbbbaccc",
			want:   false,
		},
		{
			lhs:    "a",
			rhs:    "",
			target: "a",
			want:   true,
		},
		{
			lhs:    "a",
			rhs:    "",
			target: "b",
			want:   false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, IsInterleave(tt.lhs, tt.rhs, tt.target))
	}
}
