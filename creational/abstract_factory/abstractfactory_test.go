package abstractfactory

import (
	"fmt"
	"testing"
)

func Produce(factory AbstractFactory) {
	factory.CreateProductA()
	factory.CreateProductB()
}

func Produce1() {
	// 创建抽象工厂对象
	factory := ConcreteFactory1{}
	// 通过抽象工厂来获取一系列的对象
	Produce(factory)
}

func Produce2() {
	// 创建抽象工厂对象
	factory := ConcreteFactory2{}
	// 通过抽象工厂来获取一系列的对象
	Produce(factory)
}

func TestAbstractFactory(t *testing.T) {
	Produce1()
	fmt.Println("-----------")
	Produce2()
}
