package scramble_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//goland:noinspection SpellCheckingInspection
func TestIsScramble(t *testing.T) {
	assert.Equal(t, true, IsScramble("abb", "bba"))
	assert.Equal(t, true, IsScramble("great", "rgeat"))
	assert.Equal(t, false, IsScramble("abcde", "caebd"))
	assert.Equal(t, true, IsScramble("a", "a"))
}
