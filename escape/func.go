package main

import (
	"fmt"
)

func escape(str string) {
	println("escape", str)
}

func escape2(i interface{}) {
	println("escape", i)
}

func escape3(is ...interface{}) {
	for _, v := range is {
		println(v)
	}
}

func f1() {
	str := new(string)
	*str = "test"
}

func f2() {
	str := new(string)
	*str = "test"
	escape(*str)
}

func f3() {
	str := new(string)
	*str = "test"
	escape2(*str)
}

func f4() {
	str := new(string)
	*str = "test"
	escape3(*str)
}

func f5() {
	str := new(string)
	*str = "test"
	// escape
	fmt.Println(*str)
}

func main() {
	f1()
	f2()
	f3()
	f4()
}
