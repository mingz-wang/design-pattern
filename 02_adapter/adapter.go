package adapter

// Adaptee 已经存在的接口，需要被适配
type Adaptee interface {
	SpecificRequest() string
}

// NewAdaptee 被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// adapteeImpl 被适配的类
type adapteeImpl struct{}

func (adapteeImpl) SpecificRequest() string {
	return "Adaptee Specific Request method"
}

// Target 客户端使用的接口，与特定领域相关
type Target interface {
	Request() string
}

// Adapter 适配器
type Adapter struct {
	Adaptee Adaptee
}

// NewAdapter 是 Adapter 的工厂函数
func NewAdapter(adaptee Adaptee) Adapter {
	return Adapter{Adaptee: adaptee}
}

func (a Adapter) Request() string {
	return a.Adaptee.SpecificRequest()
}
