package unique_paths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m, n, want int
	}{
		{
			m:    23,
			n:    12,
			want: 193536720,
		},
		{
			m:    3,
			n:    7,
			want: 28,
		},
		{
			m:    3,
			n:    2,
			want: 3,
		},
		{
			m:    7,
			n:    3,
			want: 28,
		},
		{
			m:    3,
			n:    3,
			want: 6,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, UniquePaths(tt.m, tt.n))
	}
}
