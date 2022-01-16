package composite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComposite(t *testing.T)  {
	file1 := &file{&meta{name: "file1",size: 10, path: "/test", content: "file1test"}}
	file2 := &file{&meta{name: "file2",size: 10, path: "/test", content: "file2test"}}
	folder1 := &folder{meta: &meta{name: "folder1", path: "/test", size: 1}}

	file1.Search("aaa")
	folder1.Search("file1")

	AddComponentToFolder(folder1, file2)
	assert.Equal(t, "/test/folder1", file2.Path())
	calFolderSize(folder1)
	assert.Equal(t, int64(11), folder1.Size())
}

func AddComponentToFolder(folder Folder, component Component)  {
	component.SetPath(folder.GetCurrentPath() + "/" + folder.GetFolderName())
	folder.AddComponent(component)
}

func calFolderSize(folder *folder) {
	size := folder.Size()
	for _,  v := range folder.components {
		size += v.Size()
	}
	folder.meta.size = size
}