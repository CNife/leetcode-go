package short_encoding_of_words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumLengthEncoding(t *testing.T) {
	tests := []struct {
		words []string
		want  int
	}{
		{
			words: []string{"time", "me", "bell"},
			want:  10,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MinimumLengthEncoding(tt.words))
	}
}
