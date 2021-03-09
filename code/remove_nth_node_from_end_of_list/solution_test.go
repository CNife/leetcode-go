package remove_nth_node_from_end_of_list

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		head *types.ListNode
		n    int
		want *types.ListNode
	}{
		{
			head: types.NewList(1, 2, 3, 4, 5),
			n:    2,
			want: types.NewList(1, 2, 3, 5),
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, RemoveNthFromEnd(tt.head, tt.n))
	}
}
