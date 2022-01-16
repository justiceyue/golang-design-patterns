package factorymethod

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	dataFactory := FileFactory{}
	dc := inputData(dataFactory, "aaa")
	fmt.Println(dc.ExportData())
}

// 底层方法并不知道实例对象及具体实现，延迟到高层方法中构造
func inputData(dataFactory DataOperatorFactory, in string) DataOperator {
	dc := dataFactory.Create()
	dc.ImportData(in)
	return dc
}
