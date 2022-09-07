package rearrange_spaces_between_words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorderSpaces(t *testing.T) {
	tests := []struct {
		text, want string
	}{
		{
			text: "  this   is  a sentence ",
			want: "this   is   a   sentence",
		},
		{
			text: " practice   makes   perfect",
			want: "practice   makes   perfect ",
		},
		{
			text: "hello   world",
			want: "hello   world",
		},
		{
			text: "  walks  udp package   into  bar a",
			want: "walks  udp  package  into  bar  a ",
		},
		{
			text: "a",
			want: "a",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, ReorderSpaces(tt.text))
	}
}
