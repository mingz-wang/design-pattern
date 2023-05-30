package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	sub := &Proxy{}

	res := sub.Do()

	if res != "before:real:after" {
		t.Fail()
	}
}
