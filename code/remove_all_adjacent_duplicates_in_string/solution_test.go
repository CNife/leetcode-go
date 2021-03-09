package remove_all_adjacent_duplicates_in_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s, want string
	}{
		{
			s:    "abbaca",
			want: "ca",
		},
		{
			s:    "",
			want: "",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, RemoveDuplicates(tt.s))
	}
}
