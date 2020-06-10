package oop

import (
	"fmt"
	"testing"
)

type Person struct {
	name        string
	age         int
	salaryMonth int

	city     string
	postCode int
}

type PersonBuilder struct {
	p *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		p: &Person{},
	}
}

func (p *PersonBuilder) Build() *Person {
	return p.p
}

func (p *PersonBuilder) Address() *PersonBuilderAddress {
	return &PersonBuilderAddress{p}
}

func (p *PersonBuilder) Data() *PersonBuilderData {
	return &PersonBuilderData{p}
}

type PersonBuilderAddress struct {
	*PersonBuilder
}

func (p *PersonBuilderAddress) Code(postCode int) *PersonBuilderAddress {
	p.p.postCode = postCode
	return p
}

func (p *PersonBuilderAddress) City(city string) *PersonBuilderAddress {
	p.p.city = city
	return p
}

type PersonBuilderData struct {
	*PersonBuilder
}

func (p *PersonBuilderData) Name(name string) *PersonBuilderData {
	p.p.name = name
	return p
}

func (p *PersonBuilderData) Age(age int) *PersonBuilderData {
	p.p.age = age
	return p
}

func (p *PersonBuilderData) Salary(salary int) *PersonBuilderData {
	p.p.salaryMonth = salary
	return p
}

func TestFactory(t *testing.T) {
	pb := NewPersonBuilder().
		Data().Name("Lukasz").Age(34).Salary(1000).
		Address().City("Krakow").Code(30485)

	p := pb.Build()

	fmt.Printf("p=%#v\n", p)
}
