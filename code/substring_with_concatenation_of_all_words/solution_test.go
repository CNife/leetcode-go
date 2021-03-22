package substring_with_concatenation_of_all_words

import (
	"testing"

	"github.com/CNife/leetcode-go/util/sort"
	"github.com/stretchr/testify/assert"
)

func TestFindSubstring(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s     string
		words []string
		want  []int
	}{
		{
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "good"},
			want:  []int{8},
		},
		{
			s:     "barfoothefoobarman",
			words: []string{"foo", "bar"},
			want:  []int{0, 9},
		},
		{
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "word"},
			want:  nil,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, sort.Ints(tt.want), FindSubstring(tt.s, tt.words))
	}
}
