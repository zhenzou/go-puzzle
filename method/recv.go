package main

import (
	"fmt"
)

type Animal interface {
	Bark()
}

type Dog struct {
}

func (d Dog) Bark() {
	fmt.Println("dog")
}

type Cat struct {
}

func (c *Cat) Bark() {
	fmt.Println("cat")
}

func Bark(a Animal) {
	a.Bark()
}

func getDog() Dog {
	return Dog{}
}

func getCat() Cat {
	return Cat{}
}

// 对于类型*T，它的方法集合是所有接收者为*T和T的方法。

func main() {
	dp := &Dog{}
	d := Dog{}
	dp.Bark() // (1)
	d.Bark()  // (2)
	Bark(dp)  // (3)
	Bark(d)   // (4)

	cp := &Cat{}
	c := Cat{}
	cp.Bark() // (5)
	c.Bark()  // (6)
	Bark(cp)  // (7)
	Bark(&c)  // (8)，这里必须取指针

	//runtime.BlockProfile()
	getDog().Bark() // (9)
	// 注意addressable，函数返回值是 not addressable
	//getCat().Bark() // (10)
}
