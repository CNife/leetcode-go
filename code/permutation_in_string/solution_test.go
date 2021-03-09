package permutation_in_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		p, s string
		want bool
	}{
		{
			p:    "adc",
			s:    "dcda",
			want: true,
		},
		{
			p:    "ab",
			s:    "eidbaooo",
			want: true,
		},
		{
			p:    "ab",
			s:    "eidboaoo",
			want: false,
		},
		{
			p:    "ab",
			s:    "a",
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, CheckInclusion(tt.p, tt.s))
	}
}
