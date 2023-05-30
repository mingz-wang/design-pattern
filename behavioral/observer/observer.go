package observer

import "fmt"

type Observer interface {
	Update(subject *Subject)
}

type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) notify() {
	for _, observer := range s.observers {
		observer.Update(s)
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{name: name}
}

func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}
