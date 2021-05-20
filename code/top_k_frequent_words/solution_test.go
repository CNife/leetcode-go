package top_k_frequent_words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		strings []string
		k       int
		want    []string
	}{
		{
			strings: []string{"i", "love", "leetcode", "i", "love", "coding"},
			k:       2,
			want:    []string{"i", "love"},
		},
		{
			strings: []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
			k:       4,
			want:    []string{"the", "is", "sunny", "day"},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, TopKFrequent(tt.strings, tt.k))
	}
}
