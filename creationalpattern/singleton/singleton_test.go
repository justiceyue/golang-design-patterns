package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingle(t *testing.T) {
	c := GetInstance()
	assert.Equal(t, 0, c.Get())
	c.Add(1)
	assert.Equal(t, 1, GetInstance().Get())
}
