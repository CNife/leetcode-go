package group_anagrams

import (
	"testing"

	"github.com/CNife/leetcode-go/util/sort"
	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strings []string
		want    [][]string
	}{
		{
			strings: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"ate", "eat", "tea"},
				{"nat", "tan"},
				{"bat"},
			},
		},
	}
	for _, tt := range tests {
		assert.Equal(t,
			sort.StringSlicesDeep(tt.want, sort.Strings),
			sort.StringSlicesDeep(GroupAnagrams(tt.strings), sort.Strings))
	}
}
