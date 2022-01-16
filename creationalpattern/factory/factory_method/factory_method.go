package factorymethod

import "fmt"

// product
type DataOperator interface {
	// 导入data,并返回是否导出成功
	ImportData(rawdata string) bool
	// 输出数据
	ExportData() string
}

type metaData struct {
	data string
}

func (m *metaData) ImportData(rawData string) bool {
	m.data = rawData
	fmt.Println("数据导入成功")
	return m.data != ""
}

type file struct {
	*metaData
}

func (f *file) ExportData() string {
	return f.data + "file"
}

type db struct {
	*metaData
}

func (d *db) ExportData() string {
	return d.data + "db"
}

// creator
type DataOperatorFactory interface {
	Create() DataOperator
}

type FileFactory struct {
}

func (FileFactory) Create() DataOperator {
	return &file{
		metaData: &metaData{},
	}
}

type DBFactory struct {
}

func (DBFactory) Create() DataOperator {
	return &db{
		metaData: &metaData{},
	}
}
