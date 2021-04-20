package implement_strstr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//goland:noinspection SpellCheckingInspection
func TestStrStr(t *testing.T) {
	assert.Equal(t, 2, StrStr("hello", "ll"))
	assert.Equal(t, -1, StrStr("aaaaa", "bba"))
	assert.Equal(t, 0, StrStr("", ""))
}
