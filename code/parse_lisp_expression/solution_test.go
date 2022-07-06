package parse_lisp_expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		expression string
		want       int
	}{
		{
			expression: "(let x 2 (mult x (let x 3 y 4 (add x y))))",
			want:       14,
		},
		{
			expression: "(let x 3 x 2 x)",
			want:       2,
		},
		{
			expression: "(let x 1 y 2 x (add x y) (add x y))",
			want:       5,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Evaluate(tt.expression))
	}
}
