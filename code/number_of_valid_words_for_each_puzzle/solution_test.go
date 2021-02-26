package number_of_valid_words_for_each_puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindNumOfValidWords(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		words, puzzles []string
		want           []int
	}{
		{
			words:   []string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"},
			puzzles: []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"},
			want:    []int{1, 1, 3, 2, 4, 0},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FindNumOfValidWords(tt.words, tt.puzzles))
	}
}
