package most_common_word

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostCommonWord(t *testing.T) {
	tests := []struct {
		paragraph string
		banned    []string
		want      string
	}{
		{

			paragraph: "Bob hit a ball, the hit BALL flew far after it was hit.",
			banned:    []string{"hit"},
			want:      "ball",
		},
		{
			paragraph: "Bob",
			banned:    []string{},
			want:      "bob",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MostCommonWord(tt.paragraph, tt.banned))
	}
}
