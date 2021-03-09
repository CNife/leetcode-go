package text_justification

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullJustify(t *testing.T) {
	tests := []struct {
		words    []string
		maxWidth int
		want     []string
	}{
		{
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			want:     []string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			want:     []string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			words: []string{
				"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a",
				"computer.", "Art", "is", "everything", "else", "we", "do",
			},
			maxWidth: 20,
			want: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FullJustify(tt.words, tt.maxWidth))
	}
}
