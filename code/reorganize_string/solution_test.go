package reorganize_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorganizeString(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		src, want string
	}{
		{
			src:  "aab",
			want: "aba",
		},
		{
			src:  "aaab",
			want: "",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, ReorganizeString(tt.src))
	}
}
