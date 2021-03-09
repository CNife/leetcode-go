package repeated_dna_sequences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		src  string
		want []string
	}{
		{
			src:  "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			want: []string{"AAAAACCCCC", "CCCCCAAAAA"},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, counter(tt.want), counter(FindRepeatedDnaSequences(tt.src)))
	}
}

func counter(ss []string) map[string]int {
	m := make(map[string]int, len(ss))
	for _, s := range ss {
		if count, exists := m[s]; exists {
			m[s] = count + 1
		} else {
			m[s] = 1
		}
	}
	return m
}
