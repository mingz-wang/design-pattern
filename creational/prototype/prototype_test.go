package prototype

import (
	"fmt"
	"testing"
)

type Client struct {
	prototype Prototype
}

func NewClient(prototype Prototype) *Client {
	return &Client{
		prototype: prototype,
	}
}

func (c *Client) operation() {
	newPrototype := c.prototype.Clone()
	fmt.Println(newPrototype)
}

func Test(t *testing.T) {
	client := NewClient(ConcretePrototype1{})
	client.operation()
	client = NewClient(ConcretePrototype2{})
	client.operation()
}
