package word_break_2

import (
	"testing"

	"github.com/CNife/leetcode-go/util"
	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []*struct {
		s     string
		words []string
		want  []string
	}{
		{
			s:     "catsanddog",
			words: []string{"cat", "cats", "and", "sand", "dog"},
			want:  []string{"cats and dog", "cat sand dog"},
		},
		{
			s:     "pineapplepenapple",
			words: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			want:  []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"},
		},
		{
			s:     "catsandog",
			words: []string{"cats", "dog", "sand", "and", "cat"},
			want:  nil,
		},
		{
			s: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
				"aaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			words: []string{
				"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa",
				"aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, util.SortStrings(tt.want),
			util.SortStrings(WordBreak(tt.s, tt.words)))
	}
}
