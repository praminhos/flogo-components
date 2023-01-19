package string

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"

	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestPadRight(t *testing.T) {
	f := &fnPadRight{}
	final, err := f.Eval("abc", 5, "a")
	assert.Nil(t, err)
	assert.Equal(t, "abcaa", final)
}
