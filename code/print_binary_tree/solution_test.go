package print_binary_tree

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestPrintTree(t *testing.T) {
	tests := []struct {
		root *types.TreeNode
		want [][]string
	}{
		{
			root: types.NewTree(1, 2),
			want: [][]string{
				{"", "1", ""},
				{"2", "", ""},
			},
		},
		{
			root: types.NewTree(1, 2, 3, -1, 4),
			want: [][]string{
				{"", "", "", "1", "", "", ""},
				{"", "2", "", "", "", "3", ""},
				{"", "", "4", "", "", "", ""},
			},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, PrintTree(tt.root))
	}
}
