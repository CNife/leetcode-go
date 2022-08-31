package validate_stack_sequences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateStackSequences(t *testing.T) {
	tests := []struct {
		pushed, popped []int
		want           bool
	}{
		{
			pushed: []int{1, 2, 3, 4, 5},
			popped: []int{4, 5, 3, 2, 1},
			want:   true,
		},
		{
			pushed: []int{1, 2, 3, 4, 5},
			popped: []int{4, 3, 5, 1, 2},
			want:   false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, ValidateStackSequences(tt.pushed, tt.popped))
	}
}
