package freedom_trail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRotateSteps(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		ring, key string
		want      int
	}{
		{
			ring: "edcba",
			key:  "abcde",
			want: 10,
		},
		{
			ring: "godding",
			key:  "gd",
			want: 4,
		},
		{
			ring: "caotmcaataijjxi",
			key:  "oatjiioicitatajtijciocjcaaxaaatmctxamacaamjjx",
			want: 137,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FindRotateSteps(tt.ring, tt.key))
	}
}
