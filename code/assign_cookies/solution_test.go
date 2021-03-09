package assign_cookies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindContentChildren(t *testing.T) {
	tests := []struct {
		g, s []int
		want int
	}{
		{
			g:    []int{1, 2, 3},
			s:    []int{1, 1},
			want: 1,
		},
		{
			g:    []int{1, 2},
			s:    []int{1, 2, 3},
			want: 2,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FindContentChildren(tt.g, tt.s))
	}
}
