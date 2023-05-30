package factorymethod

import (
	"testing"
)

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	var (
		factory OperatorFactory
	)

	factory = PlusOperatorFactory{}
	if res := compute(factory, 1, 2); res != 3 {
		t.Fatalf("want %v, but got %v", 3, res)
	}

	factory = MinusOperatorFactory{}
	if res := compute(factory, 5, 2); res != 3 {
		t.Fatalf("want %v, but got %v", 3, res)
	}
}
