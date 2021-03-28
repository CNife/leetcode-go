package rotate_list

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		head *types.ListNode
		k    int
		want *types.ListNode
	}{
		{
			head: types.NewList(1, 2, 3, 4, 5),
			k:    2,
			want: types.NewList(4, 5, 1, 2, 3),
		},
		{
			head: types.NewList(0, 1, 2),
			k:    4,
			want: types.NewList(2, 0, 1),
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, RotateRight(tt.head, tt.k))
	}
}
