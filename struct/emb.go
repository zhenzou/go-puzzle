package main

import "reflect"

type Parent struct {
	Name string
}

type Child struct {
	Parent
	Sex int
}

// 不能接受Child类型
func Println(p Parent) {
	println("p:", p.Name)
}

func Println2(i interface{}) {
	if p, ok := i.(Parent); ok {
		Println(p)
	} else {
		typ := reflect.TypeOf(i)
		println(typ.Name(), "is not type Parent")
	}
}

func NewParent() Parent {
	// 直接返回Child并不行
	return Child{
		Parent: Parent{"test"},
		Sex:    0,
	}.Parent
}

func PrintlnFields(i interface{}) {
	typ := reflect.TypeOf(i)
	if typ.Kind() == reflect.Struct {
		num := typ.NumField()
		for i := 0; i < num; i++ {
			field := typ.Field(i)
			println(field.Type.Name(), ":", field.Name)
		}
	}
}

func main() {
	child := &Child{
		Parent: Parent{Name: "test"},
		Sex:    0,
	}
	Println2(*child)
	PrintlnFields(*child)
}
