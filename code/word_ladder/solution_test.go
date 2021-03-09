package word_ladder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLadderLength(t *testing.T) {
	tests := []struct {
		beginWord, endWord string
		wordList           []string
		want               int
	}{
		{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			want:      5,
		},
		{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			want:      0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want,
			LadderLength(tt.beginWord, tt.endWord, tt.wordList))
	}
}
