package odd_even_linked_list

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestOddEvenList(t *testing.T) {
	tests := []struct {
		head, want *types.ListNode
	}{
		{
			head: types.NewList(1, 2, 3, 4, 5),
			want: types.NewList(1, 3, 5, 2, 4),
		},
		{
			head: types.NewList(2, 1, 3, 5, 6, 4, 7),
			want: types.NewList(2, 3, 6, 7, 1, 5, 4),
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, OddEvenList(tt.head))
	}
}
