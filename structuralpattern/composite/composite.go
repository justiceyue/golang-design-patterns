package composite

import (
	"fmt"
	"strings"
)

type Component interface {
	Name() string
	Size() int64
	Path() string
	Content() string

	SetPath(string)
	Search(string)
}

type meta struct {
	name    string
	size    int64
	path    string
	content string
}

func (m *meta) Name() string {
	return m.name
}

func (m *meta) Size() int64 {
	return m.size
}

func (m *meta) Path() string {
	return m.path
}

func (m *meta) Content() string {
	return m.content
}

func (m *meta) SetPath(newPath string) {
	m.path = newPath

}

type File interface {
	WriteToFile(string)
}

type file struct {
	*meta
}

func (f *file) Search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *file) WriteToFile(content string) {
	f.content = strings.Join([]string{f.content} ,content)
}

type Folder interface {
	AddComponent(Component)
	GetCurrentPath() string
	GetFolderName() string
}

type folder struct {
	components []Component
	*meta
}

func (f *folder) Search(fileName string) {
	fmt.Printf("Searching for fileName %s in folder %s\n", fileName, f.meta.name)
}

func (f *folder) AddComponent(component Component)  {
	f.components = append(f.components, component)
}

func (f *folder) GetCurrentPath() string {
	return f.path
}

func (f *folder) GetFolderName() string {
	return f.name
}