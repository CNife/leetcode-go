package clone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSlice(t *testing.T) {
	tests := [][]int{
		{1, 2, 3},
		make([]int, 0),
		nil,
	}
	for _, tt := range tests {
		assert.Equal(t, tt, IntSlice(tt))
	}
}

func TestStringSlice(t *testing.T) {
	tests := [][]string{
		{"abc", "efg", ""},
		nil,
	}
	for _, tt := range tests {
		assert.Equal(t, tt, StringSlice(tt))
	}
}

func TestIntMatrix(t *testing.T) {
	tests := [][][]int{
		nil,
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{{}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt, IntMatrix(tt))
	}
}

func TestStringMatrix(t *testing.T) {
	tests := [][][]string{
		nil,
		{
			{"a", "b", "c"},
			{"de", "fg", "hi"},
		},
		{{}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt, StringMatrix(tt))
	}
}
