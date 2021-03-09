package palindrome_linked_list

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		head *types.ListNode
		want bool
	}{
		{
			head: types.NewList(1, 2),
			want: false,
		},
		{
			head: types.NewList(1, 2, 2, 1),
			want: true,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, IsPalindrome(tt.head))
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	list := makePalindromeList(1_000_000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IsPalindrome(list)
	}
}

func makePalindromeList(size int) *types.ListNode {
	values := make([]int, size)
	for i := 0; i < size/2; i++ {
		values[i] = i
	}
	for i := size / 2; i < size; i++ {
		values[i] = size - 1 - i
	}
	return types.NewList(values...)
}
