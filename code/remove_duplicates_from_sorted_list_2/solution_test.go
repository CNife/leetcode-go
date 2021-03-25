package remove_duplicates_from_sorted_list_2

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
			head: types.NewList(1, 2, 3, 3, 4, 4, 5),
			want: types.NewList(1, 2, 5),
		},
		{
			head: types.NewList(1, 1, 1, 2, 3),
			want: types.NewList(2, 3),
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, DeleteDuplicates(tt.head))
	}
}
