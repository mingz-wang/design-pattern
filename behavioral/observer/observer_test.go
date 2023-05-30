package observer

import "testing"

func TestObserver(t *testing.T) {
	subject := NewSubject()
	reader1 := NewReader("Alice")
	reader2 := NewReader("Bob")
	reader3 := NewReader("Charlie")
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

	subject.UpdateContext("observer mode")
}
