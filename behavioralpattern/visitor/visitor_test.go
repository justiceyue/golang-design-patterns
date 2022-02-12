package visitor

import (
	"testing"
)

func TestVisitor(t *testing.T) {
	fl := fileList{
		&PDF{"p1"},
		&TXT{"t1"},
		&TXT{"t2"},
	}
	e := &ExactFile{}
	c := &CompressFile{}
	fl.Iter(func(f File) {
		//二次转发
		//File->pdf or txt -> visitor
		f.Accept(e)
	})
	fl.Iter(func(f File) {
		f.Accept(c)
	})
}

type fileList []File

func (fl fileList) Iter(f func(f File)) {
	for _, file := range fl {
		f(file)
	}
}
