package distinct_subsequences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumDistinct(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s, t string
		want int
	}{
		{
			s:    "rabbbit",
			t:    "rabbit",
			want: 3,
		},
		{
			s:    "babgbag",
			t:    "bag",
			want: 5,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, NumDistinct(tt.s, tt.t))
	}
}
