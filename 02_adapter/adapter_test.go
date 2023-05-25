package adapter

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	// 创建需要被适配的对象
	adaptee := NewAdaptee()
	// 创建客户端需要调用的接口对象
	target := NewAdapter(adaptee)
	// 请求处理
	res := target.Request()
	fmt.Println(res)
}
