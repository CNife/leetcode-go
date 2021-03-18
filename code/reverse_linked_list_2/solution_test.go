package reverse_linked_list_2

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		head        *types.ListNode
		left, right int
		want        *types.ListNode
	}{
		{
			head:  types.NewList(1, 2, 3, 4, 5, 6),
			left:  3,
			right: 5,
			want:  types.NewList(1, 2, 5, 4, 3, 6),
		},
		{
			head:  types.NewList(1, 2, 3),
			left:  1,
			right: 3,
			want:  types.NewList(3, 2, 1),
		},
		{
			head:  nil,
			left:  1,
			right: 1,
			want:  nil,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, ReverseBetween(tt.head, tt.left, tt.right))
	}
}
