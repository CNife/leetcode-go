package palindrome_partitioning

import (
	"testing"

	"github.com/CNife/leetcode-go/util/sort"
	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		s    string
		want [][]string
	}{
		{
			s: "aab",
			want: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			s: "a",
			want: [][]string{
				{"a"},
			},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, sort.StringSlices(tt.want),
			sort.StringSlices(Partition(tt.s)))
	}
}
