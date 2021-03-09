package group_anagrams

import (
	"testing"

	"github.com/CNife/leetcode-go/util"
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
			util.SortStringSlicesDeep(tt.want, util.SortStrings),
			util.SortStringSlicesDeep(GroupAnagrams(tt.strings), util.SortStrings))
	}
}
