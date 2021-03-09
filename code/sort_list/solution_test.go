package sort_list

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestSortList(t *testing.T) {
	tests := []struct {
		head, want *types.ListNode
	}{
		{
			head: types.NewList(4, 2, 1, 3),
			want: types.NewList(1, 2, 3, 4),
		},
		{
			head: types.NewList(-1, 5, 3, 4, 0),
			want: types.NewList(-1, 0, 3, 4, 5),
		},
		{
			head: nil,
			want: nil,
		},
		{
			head: types.NewList(0),
			want: types.NewList(0),
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, SortList(tt.head))
	}
}
