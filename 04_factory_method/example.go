package factorymethod

import "fmt"

// exportFileApi 导出的文件对象的接口
type exportFileApi interface {
	export(data string) bool
}

// exportTxtFile 导出成文本文件格式的对象
type exportTxtFile struct{}

func (exportTxtFile) export(data string) bool {
	fmt.Println("导出数据 <" + data + "> 到文本文件中")
	return true
}

// exportDB 导出成数据库备份文件的对象
type exportDB struct{}

func (exportDB) export(data string) bool {
	fmt.Println("导出数据 <" + data + "> 到数据库备份文件")
	return true
}

// ExportOperation 实现导出数据的业务功能对象，被封装的实际类接口
type ExportOperation interface {
	Export(data string) bool
}

// exportOperationBase 接口实现的基类，封装公共方法
type exportOperationBase struct {
	api exportFileApi
}

func (e exportOperationBase) Export(data string) bool {
	return e.api.export(data)
}

// exportTxtFileOperation 导出成文本文件格式
type exportTxtFileOperation struct {
	*exportOperationBase
}

// exportDBOperation 导出成数据库文件
type exportDBOperation struct {
	*exportOperationBase
}

// NewExportTxtFileOperation 工厂方法
func NewExportTxtFileOperation() ExportOperation {
	return &exportTxtFileOperation{
		exportOperationBase: &exportOperationBase{api: exportTxtFile{}},
	}
}

// NewExportDBOperation 工厂方法
func NewExportDBOperation() ExportOperation {
	return &exportDBOperation{exportOperationBase: &exportOperationBase{api: exportDB{}}}
}
