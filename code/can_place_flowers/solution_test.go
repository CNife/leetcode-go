package can_place_flowers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		flowerbed []int
		n         int
		want      bool
	}{
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			want:      true,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			want:      false,
		},
		{
			flowerbed: []int{1},
			n:         0,
			want:      true,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, CanPlaceFlowers(tt.flowerbed, tt.n))
	}
}
