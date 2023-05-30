package factorymethod

// Operator 被封装的实际接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// 是 Operator 接口实现的基类，封装公共方法
type operatorBase struct {
	a int
	b int
}

func (o *operatorBase) SetA(a int) {
	o.a = a
}

func (o *operatorBase) SetB(b int) {
	o.b = b
}

// PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	*operatorBase
}

// Result 获取结果
func (plus PlusOperator) Result() int {
	return plus.a + plus.b
}

// PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct{}

func (p PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		operatorBase: &operatorBase{},
	}
}

// MinusOperator Operator 的实际减法实现
type MinusOperator struct {
	*operatorBase
}

// Result 获取结果
func (plus MinusOperator) Result() int {
	return plus.a - plus.b
}

// MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct{}

func (p MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		operatorBase: &operatorBase{},
	}
}
