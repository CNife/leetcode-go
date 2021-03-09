package rectangle_area

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeArea(t *testing.T) {
	tests := []struct {
		A, B, C, D, E, F, G, H, want int
	}{
		{
			A:    -3,
			B:    0,
			C:    3,
			D:    4,
			E:    0,
			F:    -1,
			G:    9,
			H:    2,
			want: 45,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want,
			ComputeArea(tt.A, tt.B, tt.C, tt.D, tt.E, tt.F, tt.G, tt.H))
	}
}
