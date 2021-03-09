package remove_k_digits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveKDigits(t *testing.T) {
	tests := []struct {
		num  string
		k    int
		want string
	}{
		{
			num:  "1432219",
			k:    3,
			want: "1219",
		},
		{
			num:  "10200",
			k:    1,
			want: "200",
		},
		{
			num:  "10",
			k:    1,
			want: "0",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, RemoveKDigits(tt.num, tt.k))
	}
}
