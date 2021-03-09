package reorder_list

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		head, want *types.ListNode
	}{
		{
			head: types.NewList(1, 2, 3, 4),
			want: types.NewList(1, 4, 2, 3),
		},
		{
			head: types.NewList(1, 2, 3, 4, 5),
			want: types.NewList(1, 5, 2, 4, 3),
		},
	}
	for _, tt := range tests {
		ReorderList(tt.head)
		assert.Equal(t, tt.want, tt.head)
	}
}
