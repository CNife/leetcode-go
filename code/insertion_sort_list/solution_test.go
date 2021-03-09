package insertion_sort_list

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestInsertionSortList(t *testing.T) {
	tests := []struct {
		head, want *types.ListNode
	}{
		{
			head: nil,
			want: nil,
		},
		{
			head: types.NewList(1),
			want: types.NewList(1),
		},
		{
			head: types.NewList(4, 2, 1, 3),
			want: types.NewList(1, 2, 3, 4),
		},
		{
			head: types.NewList(-1, 5, 3, 4, 0),
			want: types.NewList(-1, 0, 3, 4, 5),
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, InsertionSortList(tt.head))
	}
}
