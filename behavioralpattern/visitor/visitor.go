package visitor

import "fmt"

/*
	给一系列对象透明的添加新功能
	通过预留调用通路和两次分发技术来实现
*/

//file接口只提供一个简单的读文件操作，通过Accpet可实现功能的扩展
type File interface {
	Read()
	//预留调用通路
	Accept(v Visitor)
}

type PDF struct {
	FileName string
}

func (p *PDF) Read() {
	fmt.Println(p.FileName)
}

func (p *PDF) Accept(v Visitor) {
	v.VisitForPDF(p)
}

type TXT struct {
	FileName string
}

func (T *TXT) Read() {
	fmt.Println(T.FileName)
}

func (T *TXT) Accept(v Visitor) {
	v.VisitForTXT(T)
}

//虽然访问者模式可以很好的扩展功能，但是也导致了接口不稳定
type Visitor interface {
	VisitForPDF(p *PDF)
	VisitForTXT(p *TXT)
}

//解压文件操作
type ExactFile struct{}

func (e *ExactFile) VisitForPDF(p *PDF) {
	fmt.Printf("解压pdf name:%s\n", p.FileName)
}

func (e *ExactFile) VisitForTXT(t *TXT) {
	fmt.Printf("解压txt name:%s\n", t.FileName)
}

//压缩文件操作

type CompressFile struct{}

func (c *CompressFile) VisitForPDF(p *PDF) {
	fmt.Printf("压缩pdf name:%s\n", p.FileName)
}

func (c *CompressFile) VisitForTXT(t *TXT) {
	fmt.Printf("压缩txt name:%s\n", t.FileName)
}
