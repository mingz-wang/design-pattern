package prototype

// Prototype 声明一个克隆自身的接口
type Prototype interface {
	Clone() Prototype
}

// ConcretePrototype1 克隆的具体的实现对象1
type ConcretePrototype1 struct {
	name string
}

func (ConcretePrototype1) Clone() Prototype {
	prototype := ConcretePrototype1{name: "concrete prototype 1"}
	return prototype
}

// ConcretePrototype2 克隆的具体的实现对象2
type ConcretePrototype2 struct {
	name string
}

func (ConcretePrototype2) Clone() Prototype {
	prototype := ConcretePrototype2{name: "concrete prototype 2"}
	return prototype
}
