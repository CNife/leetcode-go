package lemonade_change

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLemonadeChange(t *testing.T) {
	tests := []struct {
		bills []int
		want  bool
	}{
		{
			bills: []int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5},
			want:  true,
		},
		{
			bills: []int{5, 5, 5, 10, 20},
			want:  true,
		},
		{
			bills: []int{5, 5, 10},
			want:  true,
		},
		{
			bills: []int{10, 10},
			want:  false,
		},
		{
			bills: []int{5, 5, 10, 10, 20},
			want:  false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, LemonadeChange(tt.bills))
	}
}
