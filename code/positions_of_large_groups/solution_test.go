package positions_of_large_groups

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargeGroupPosition(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s    string
		want [][]int
	}{
		{
			s:    "abbxxxxzzy",
			want: [][]int{{3, 6}},
		},
		{
			s:    "abc",
			want: nil,
		},
		{
			s:    "abcdddeeeeaabbbcd",
			want: [][]int{{3, 5}, {6, 9}, {12, 14}},
		},
		{
			s:    "aba",
			want: nil,
		},
		{
			s:    "aaa",
			want: [][]int{{0, 2}},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, LargeGroupPosition(tt.s))
	}
}
