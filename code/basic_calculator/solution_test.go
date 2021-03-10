package basic_calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "1 + 1",
			want: 2,
		},
		{
			s:    " 2-1 + 2 ",
			want: 3,
		},
		{
			s:    "(1+(4+5+2)-3)+(6+8)",
			want: 23,
		},
		{
			s:    "100",
			want: 100,
		},
		{
			s:    "(1)",
			want: 1,
		},
		{
			s:    "-2+ 1",
			want: -1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Calculate(tt.s))
	}
}
