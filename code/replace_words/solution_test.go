package replace_words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceWords(t *testing.T) {
	tests := []struct {
		dictionary     []string
		sentence, want string
	}{
		{
			dictionary: []string{"cat", "bat", "rat"},
			sentence:   "the cattle was rattled by the battery",
			want:       "the cat was rat by the bat",
		}, {
			dictionary: []string{"a", "b", "c"},
			sentence:   "aadsfasf absbs bbab cadsfafs",
			want:       "a a b c",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, ReplaceWords(tt.dictionary, tt.sentence))
	}
}
