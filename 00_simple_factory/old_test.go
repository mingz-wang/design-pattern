package simplefactory

// import (
// 	"fmt"
// 	"testing"
// )
//
// // Api 某个接口
// type Api interface {
// 	test(s string)
// }
//
// // Impl 对接口的实现
// type Impl struct{}
//
// func (Impl) test(s string) {
// 	fmt.Printf("Now In Impl. The input s == %v\n", s)
// }
//
// // TestUse 测试使用 Api 接口
// func TestUse(t *testing.T) {
// 	var api Api = Impl{}
// 	api.test("这就是个测试，啦啦啦。")
// }
