package simplefactory

import "fmt"

// Api 接口的定义，该接口可以通过简单工厂创建
type Api interface {
	Operation(s string)
}

// ImplA 接口的实现
type ImplA struct{}

func (ImplA) Operation(s string) {
	fmt.Printf("ImplA s == %v\n", s)
}

// ImplB 接口的实现
type ImplB struct{}

func (ImplB) Operation(s string) {
	fmt.Printf("ImplB s == %v\n", s)
}

// CreateApi 工厂函数，用来创建 Api 对象
func CreateApi(condition int) Api {
	// 应该根据某些条件去选择究竟创建哪一个具体的实现对象,
	// 这些条件可以从外部传入,也可以从其他途径来获取。
	// 如果只有一个实现,可以省略条件,因为没有选择的必要。

	// 示意使用条件
	var api Api
	switch condition {
	case 1:
		api = ImplA{}
	case 2:
		api = ImplB{}
	}
	return api
}
