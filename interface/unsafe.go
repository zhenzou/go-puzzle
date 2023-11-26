package main

import (
	"reflect"
)

type Runnable interface {
	Run()
}

type Container struct {
	Runnable Runnable
}

func main() {
	container := Container{}

	typ := reflect.TypeOf(container).Field(0).Type
	println(typ.String())

	// typ := reflect.TypeOf(container.Runnable)
	// println(typ.String())

	val := reflect.ValueOf(container).Field(0)
	// val := reflect.ValueOf(container.Runnable)
	println(val.IsNil())
	println(val.IsZero())
	println(val.IsValid())

	println(container.Runnable == nil)

	println(val.UnsafePointer() == nil)
}
