package remove_duplicates_from_sorted_list

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		head, want *types.ListNode
	}{
		{
			head: types.NewList(1, 1, 2),
			want: types.NewList(1, 2),
		},
		{
			head: types.NewList(1, 1, 2, 3, 3),
			want: types.NewList(1, 2, 3),
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, DeleteDuplicates(tt.head))
	}
}
