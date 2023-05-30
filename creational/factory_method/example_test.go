package factorymethod

import "testing"

func Test(t *testing.T) {
	// 创建需要使用的 Creator 对象
	operate := NewExportTxtFileOperation()
	// 调用输出数据的功能方法
	operate.Export("测试数据")
	// 创建需要使用的 Creator 对象
	operate = NewExportDBOperation()
	// 调用输出数据的功能方法
	operate.Export("测试数据")
}
