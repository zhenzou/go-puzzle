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

// 结论：
// reflect.New return ptr to type
func main() {

	typ := reflect.TypeOf(&Person{})

	value := reflect.New(typ) // value type is ptr of ptr

	println(value.IsNil()) // false

	p, ok := value.Interface().(Wrapper)
	if ok {
		// not ok
		println("name:", p.NativeType())
	}

	p, ok = value.Elem().Interface().(Wrapper)
	if ok {
		// ok
		println("name:", p.NativeType().String())
	}
}
