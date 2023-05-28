package prototype

import (
	"fmt"
	"testing"
)

var pm *PrototypeManager

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

type Type2 struct{}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

func TestClone(t *testing.T) {
	t1 := pm.Get("t1")

	t2 := t1.Clone()

	if t1 == t2 {
		t.Fatal("Error! Get clone not working")
	}
}

func TestCloneFromManager(t *testing.T) {
	pm := NewPrototypeManager()

	t1 := &Type1{name: "type1"}

	pm.Set("t1", t1)

	t2 := pm.Get("t1").Clone()

	fmt.Println(t1, t2)

	if t1.name != "type1" {
		t.Fatal("Error!")
	}
}
