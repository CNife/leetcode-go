package check_if_a_word_occurs_as_a_prefix_of_any_word_in_a_sentence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrefixOfWord(t *testing.T) {
	tests := []struct {
		sentence, searchWord string
		want                 int
	}{
		{
			sentence:   "i love eating burger",
			searchWord: "burg",
			want:       4,
		},
		{
			sentence:   "this problem is an easy problem",
			searchWord: "pro",
			want:       2,
		},
		{
			sentence:   "i am tired",
			searchWord: "you",
			want:       -1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, IsPrefixOfWord(tt.sentence, tt.searchWord))
	}
}
