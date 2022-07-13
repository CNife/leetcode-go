package asteroid_collision

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		asteroids, want []int
	}{
		{
			asteroids: []int{5, 10, -5},
			want:      []int{5, 10},
		},
		{
			asteroids: []int{8, -8},
			want:      []int{},
		},
		{
			asteroids: []int{10, 2, -5},
			want:      []int{10},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, AsteroidCollision(tt.asteroids))
	}
}
