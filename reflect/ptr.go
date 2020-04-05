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

func main() {

	typ := reflect.TypeOf(&Person{})

	value := reflect.New(typ) // value type is ptr of ptr

	println(value.IsNil()) // false

	p, ok := value.Interface().(Wrapper)
	if ok {
		println("name:", p.NativeType())
	}
}
