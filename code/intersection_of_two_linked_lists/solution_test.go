package intersection_of_two_linked_lists

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGetIntersectionNode(t *testing.T) {
	tests := []struct {
		listA, listB, listC *types.ListNode
	}{
		{
			listA: types.NewList(4, 1),
			listB: types.NewList(5, 0, 1),
			listC: types.NewList(8, 4, 5),
		},
		{
			listA: types.NewList(0, 9, 1),
			listB: types.NewList(3),
			listC: types.NewList(2, 4),
		},
		{
			listA: types.NewList(2, 6, 4),
			listB: types.NewList(1, 5),
			listC: nil,
		},
	}
	for _, tt := range tests {
		tt.listA.Last().Next = tt.listC
		tt.listB.Last().Next = tt.listC
		assert.Equal(t, tt.listC, GetIntersectionNode(tt.listA, tt.listB))
	}
}
