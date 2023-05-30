package proxy

type Subject interface {
	DO() string
}

type RealSubject struct{}

func (r *RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var res string

	// 在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等
	res += "before:"

	// 	调用真实对象
	res += p.real.Do()

	// 调用之后的操作
	res += ":after"

	return res
}
