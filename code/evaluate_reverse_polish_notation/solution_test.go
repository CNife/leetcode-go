package evaluate_reverse_polish_notation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		tokens []string
		want   int
	}{
		{
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, EvalRPN(tt.tokens))
	}
}
