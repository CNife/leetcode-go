package add_binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a, b, want string
	}{
		{
			a:    "11",
			b:    "1",
			want: "100",
		},
		{
			a:    "1010",
			b:    "1011",
			want: "10101",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, AddBinary(tt.a, tt.b))
	}
}
