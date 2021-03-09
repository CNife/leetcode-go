package longest_common_prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//noinspection SpellCheckingInspection
func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strings []string
		want    string
	}{
		{
			strings: nil,
			want:    "",
		},
		{
			strings: []string{"flower", "flow", "flight"},
			want:    "fl",
		},
		{
			strings: []string{"dog", "racecar", "car"},
			want:    "",
		},
		{
			strings: []string{"aa", "a"},
			want:    "a",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, LongestCommonPrefix(tt.strings))
	}
}
