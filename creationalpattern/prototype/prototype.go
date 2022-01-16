package prototype

/*
以文件系统来举例说明
file和folder都是inode
folder中包含file和foler
*/
type Inode interface {
	Name() string
	Clone() Inode
	Content() string
}

type file struct {
	name    string
	content string
}

func (f *file) Name() string {
	return f.name
}

func (f *file) Content() string {
	return f.content
}

func (f *file) Clone() Inode {
	return &file{
		name:    f.name + "_clone",
		content: f.content,
	}
}

func NewFile(name string, content string) *file {
	return &file{name: name, content: content}
}

type folder struct {
	name string
	// content :file name + folder name
	children []Inode
}

func (f *folder) Name() string {
	return f.name
}

func (f *folder) Content() string {
	var tempContent string
	for _, v := range f.children {
		tempContent += v.Name()
	}
	return tempContent
}

func (f *folder) Clone() Inode {
	tempFolder := &folder{name: f.name + "_clone"}
	for _, v := range f.children {
		tempFolder.children = append(tempFolder.children, v.Clone())
	}
	return tempFolder
}

func NewFolder(name string, children ...Inode) *folder {
	return &folder{
		name:     name,
		children: children,
	}
}
