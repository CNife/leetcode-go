package palindrome_partitioning_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCut(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "aab",
			want: 1,
		},
		{
			s:    "a",
			want: 0,
		},
		{
			s:    "ab",
			want: 1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MinCut(tt.s))
	}
}
