package main

import "reflect"

type Wrapper interface {
	NativeType() reflect.Type
}

type Person struct {
	Name string
}

func (person *Person) Fuck() string {
	return "fuck"
}

func (person *Person) NativeType() reflect.Type {
	return reflect.TypeOf("fuck")
}

var wrapperType = reflect.TypeOf((*Wrapper)(nil)).Elem()

func main() {
	// 指针实现Wrapper
	p := Person{}

	println(reflect.TypeOf(p).Implements(wrapperType))
	println(reflect.PtrTo(reflect.TypeOf(p)).Implements(wrapperType))
}
