package oop

import (
	"fmt"
	"testing"
)

type LazyPersonBuilder struct {
	actions []actionFn
}

func NewLazyPersonBuilder() *LazyPersonBuilder {
	return &LazyPersonBuilder{}
}

type actionFn func(p *Person)

func (pb *LazyPersonBuilder) Build() *Person {
	p := &Person{}

	for _, action := range pb.actions {
		action(p)
	}

	return p
}

func (pb *LazyPersonBuilder) Name(name string) *LazyPersonBuilder {
	pb.actions = append(pb.actions, func(p *Person) {
		p.name = name
	})
	return pb
}

func (pb *LazyPersonBuilder) Age(age int) *LazyPersonBuilder {
	pb.actions = append(pb.actions, func(p *Person) {
		p.age = age
	})
	return pb
}

func (pb *LazyPersonBuilder) Salary(salary int) *LazyPersonBuilder {
	pb.actions = append(pb.actions, func(p *Person) {
		p.salaryMonth = salary
	})
	return pb
}

func TestBuilderLazy(t *testing.T) {
	pb := NewLazyPersonBuilder().Name("Lukasz").Age(34).Salary(10)
	p := pb.Build()

	fmt.Printf("p=%#v\n", p)
}
