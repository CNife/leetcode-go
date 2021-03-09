package find_common_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonChars(t *testing.T) {
	tests := []struct {
		strings []string
		want    []string
	}{
		{
			strings: []string{"bella", "label", "roller"},
			want:    []string{"e", "l", "l"},
		},
		{
			strings: []string{"cool", "lock", "cook"},
			want:    []string{"c", "o"},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, CommonChars(tt.strings))
	}
}
