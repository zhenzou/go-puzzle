package main

import "reflect"

func Set[T any](s string) T {
	t := new(T)

	// will panic
	// can only set on Ptr or Struct field
	v := reflect.ValueOf(t)
	v.SetString(s)

	// v := reflect.ValueOf(t)
	// v.Elem().SetString(s)

	return *t
}

func main() {
	set := Set[string]("test")

	println(set)
}
