package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrototype(t *testing.T) {
	file := NewFile("file", "aaa")
	fileClone := file.Clone()
	assert.Equal(t, file.Content(), fileClone.Content())
	assert.NotEqual(t, file.Name(), fileClone.Name())
	t.Log(fileClone.Name())

	folder := NewFolder("folder", file)
	folderClone := folder.Clone()
	assert.NotEqual(t, folder.Name(), folderClone.Name())
	// file也被clone了，符合预期
	assert.Equal(t, "file_clone", folderClone.Content())
}
