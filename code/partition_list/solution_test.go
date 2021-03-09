package partition_list

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		head *types.ListNode
		x    int
		want *types.ListNode
	}{
		{
			head: types.NewList(1, 4, 3, 2, 5, 2),
			x:    3,
			want: types.NewList(1, 2, 2, 4, 3, 5),
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Partition(tt.head, tt.x))
	}
}
