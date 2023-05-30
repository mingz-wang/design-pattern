package simplefactory

import "testing"

// 使用示例
func TestSimpleFactory(t *testing.T) {
	// 通过简单工厂来获取接口对象
	api := CreateApi(1)
	api.Operation("正在使用简单工厂")
}
