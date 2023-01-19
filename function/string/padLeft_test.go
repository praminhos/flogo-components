package string

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"

	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestPadLeft(t *testing.T) {
	f := &fnPadLeft{}
	final, err := f.Eval("abc", 5, "a")
	assert.Nil(t, err)
	assert.Equal(t, "aaabc", final)
}
