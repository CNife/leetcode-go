package prefix_and_suffix_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordFilter(t *testing.T) {
	wf := Constructor([]string{"apple", "apbcle"})
	assert.Equal(t, 0, wf.F("app", "le"))
	assert.Equal(t, 1, wf.F("ap", "le"))
	assert.Equal(t, -1, wf.F("abc", "le"))

	wf = Constructor([]string{"c", "c", "i", "i"})
	assert.Equal(t, 1, wf.F("c", "c"))
	assert.Equal(t, 3, wf.F("i", "i"))
}
