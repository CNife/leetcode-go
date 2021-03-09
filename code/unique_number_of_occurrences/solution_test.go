package unique_number_of_occurrences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		array []int
		want  bool
	}{
		{
			array: []int{1, 2, 2, 1, 1, 3},
			want:  true,
		},
		{
			array: []int{1, 2},
			want:  false,
		},
		{
			array: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			want:  true,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, UniqueOccurrences(tt.array))
	}
}
