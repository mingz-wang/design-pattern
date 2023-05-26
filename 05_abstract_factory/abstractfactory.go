package abstractfactory

import "fmt"

// AbstractFactory 抽象工厂接口，声明创建抽象产品的操作
type AbstractFactory interface {
	// CreateProductA 创建抽象产品A
	CreateProductA() AbstractProductA
	// CreateProductB 创建抽象产品B
	CreateProductB() AbstractProductB
}

// AbstractProductA 抽象产品 A 的接口
type AbstractProductA interface {
	// 	定义抽象产品 A 的相关操作
}

// AbstractProductB 抽象产品 B 的接口
type AbstractProductB interface {
	// 	定义抽象产品 B 的相关操作
}

// ProductA1 产品 A 的实现 1
type ProductA1 struct{}

// ProductA2 产品 A 的实现 2
type ProductA2 struct{}

// ProductB1 产品 B 的实现 1
type ProductB1 struct{}

// ProductB2 产品 B 的实现 2
type ProductB2 struct{}

// ConcreteFactory1 具体的工厂实现 1
type ConcreteFactory1 struct{}

func (ConcreteFactory1) CreateProductA() AbstractProductA {
	fmt.Println("Product A1 is produced by ConcreteFactory1")
	return &ProductA1{}
}

func (ConcreteFactory1) CreateProductB() AbstractProductB {
	fmt.Println("Product B1 is produced by ConcreteFactory1")
	return &ProductB1{}
}

// ConcreteFactory2 具体的工厂实现 2
type ConcreteFactory2 struct{}

func (ConcreteFactory2) CreateProductA() AbstractProductA {
	fmt.Println("Product A2 is produced by ConcreteFactory2")
	return &ProductA2{}
}

func (ConcreteFactory2) CreateProductB() AbstractProductB {
	fmt.Println("Product B2 is produced by ConcreteFactory2")
	return &ProductB2{}
}
