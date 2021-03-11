package basic_calculator_2

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
			s:    "3+2*2",
			want: 7,
		},
		{
			s:    " 3/2 ",
			want: 1,
		},
		{
			s:    " 3+5 / 2 ",
			want: 5,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Calculate(tt.s))
	}
}
