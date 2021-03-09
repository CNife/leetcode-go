package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCandy(t *testing.T) {
	tests := []struct {
		ratings []int
		want    int
	}{
		{
			ratings: []int{1, 0, 2},
			want:    5,
		},
		{
			ratings: []int{1, 2, 2},
			want:    4,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Candy(tt.ratings))
	}
}
