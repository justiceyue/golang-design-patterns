package interpreter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterpreter(t *testing.T) {
	p := &Parser{}
	p.Parse("1 + 2 + 3 - 4 + 5 - 6")
	res := p.Result().Interpret()
	assert.Equal(t, 1, res)
}
