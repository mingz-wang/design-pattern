package builder

//
// // Builder 生成器接口，定义创建一个产品对象所需的各个部件的操作
// type Builder interface {
// 	BuildPart()
// }
//
// // ConcreteBuilder 具体的生成器实现
// type ConcreteBuilder struct {
// 	ResultProduct Product
// }
//
// func (ConcreteBuilder) BuildPart() {
// 	// 	构建某个部件的功能处理
// }
//
// // Product 被构建的产品的对象接口
// type Product interface {
// 	// 	定义产品的操作
// }
//
// // Director 指导者，指导使用生成器的接口来构建产品的对象
// type Director struct {
// 	builder Builder
// }
//
// // NewDirector 构造函数，传入生成器对象
// func NewDirector(builder Builder) Director {
// 	return Director{builder: builder}
// }
//
// func (d Director) Construct() {
// 	d.builder.BuildPart()
// }
