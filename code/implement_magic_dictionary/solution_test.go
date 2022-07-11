package implement_magic_dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMagicDictionary(t *testing.T) {
	md := Constructor()
	md.BuildDict([]string{"hello", "leetcode"})
	assert.False(t, md.Search("hello"))
	assert.True(t, md.Search("hhllo"))
	assert.False(t, md.Search("hell"))
	assert.False(t, md.Search("leetcoded"))

	md = Constructor()
	md.BuildDict([]string{"hello", "hallo", "leetcode"})
	assert.True(t, md.Search("hello"))
	assert.True(t, md.Search("hhllo"))
	assert.False(t, md.Search("hell"))
	assert.False(t, md.Search("leetcoded"))
}
