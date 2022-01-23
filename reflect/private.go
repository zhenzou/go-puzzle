package main

import "reflect"

type Person struct {
	name      string
	FirstName string
	LastName  string
}

// private field readonly by reflect
func main() {

	p := Person{
		name:      "name",
		FirstName: "first",
		LastName:  "last",
	}

	rv := reflect.ValueOf(p)
	println(rv.FieldByName("FirstName").String())

	println(rv.FieldByName("name").String())

	// false
	println(rv.FieldByName("name").CanSet())
	//rv.FieldByName("name").SetString("test")
}
