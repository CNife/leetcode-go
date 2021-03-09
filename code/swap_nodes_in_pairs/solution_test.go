package swap_nodes_in_pairs

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		head, want *types.ListNode
	}{
		{
			head: types.NewList(1, 2, 3, 4),
			want: types.NewList(2, 1, 4, 3),
		},
		{
			head: types.NewList(),
			want: types.NewList(),
		},
		{
			head: types.NewList(1),
			want: types.NewList(1),
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, SwapPairs(tt.head))
	}
}
