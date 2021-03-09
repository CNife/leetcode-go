package valid_palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			s:    "race a car",
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, IsPalindrome(tt.s))
	}
}
