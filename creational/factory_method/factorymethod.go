package factorymethod

import "fmt"

// Product 工厂方法所创建的对象的接口
type Product interface{}

// ConcreteProduct 具体的 Product 对象
type ConcreteProduct struct{}

// Creator 创建器
type Creator interface {
	factoryMethod() Product
	someOperation()
}

// ConcreteCreator 具体的创建器实现对象
type ConcreteCreator struct{}

// 工厂方法，返回具体的 Product 对象
func (ConcreteCreator) factoryMethod() Product {
	return ConcreteProduct{}
}

func (ConcreteCreator) someOperation() {
	fmt.Println("some operation")
}
