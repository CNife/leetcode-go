package russian_doll_envelopes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxEnvelopes(t *testing.T) {
	tests := []struct {
		envelopes [][]int
		want      int
	}{
		{
			envelopes: [][]int{
				{5, 4},
				{6, 4},
				{6, 7},
				{2, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MaxEnvelopes(tt.envelopes))
	}
}
