package string_matching_in_an_array

import (
	"testing"

	"github.com/CNife/leetcode-go/util/sort"
	"github.com/stretchr/testify/assert"
)

func TestStringMatching(t *testing.T) {
	tests := []struct {
		words, want []string
	}{
		{
			words: []string{"mass", "as", "hero", "superhero"},
			want:  []string{"as", "hero"},
		},
		{
			words: []string{"leetcode", "et", "code"},
			want:  []string{"et", "code"},
		},
		{
			words: []string{"blue", "green", "bu"},
			want:  []string{},
		},
		{
			words: []string{"leetcoder", "leetcode", "od", "hamlet", "am"},
			want:  []string{"leetcode", "od", "am"},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, sort.Strings(tt.want), sort.Strings(StringMatching(tt.words)))
	}
}
