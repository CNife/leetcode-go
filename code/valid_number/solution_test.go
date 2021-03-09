package valid_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "0",
			want:  true,
		},
		{
			input: " 0.1 ",
			want:  true,
		},
		{
			input: "abc",
			want:  false,
		},
		{
			input: "1 a",
			want:  false,
		},
		{
			input: "2e10",
			want:  true,
		},
		{
			input: " -90e3 ",
			want:  true,
		},
		{
			input: " 1e",
			want:  false,
		},
		{
			input: "e3",
			want:  false,
		},
		{
			input: " 6e-1",
			want:  true,
		},
		{
			input: " 99e2.5",
			want:  false,
		},
		{
			input: "53.5e93",
			want:  true,
		},
		{
			input: " --6 ",
			want:  false,
		},
		{
			input: "-+3",
			want:  false,
		},
		{
			input: "95a54e53",
			want:  false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, IsNumber(tt.input))
	}
}
