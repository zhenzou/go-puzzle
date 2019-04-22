package main

import (
	"go-puzzle/util"
)

type Super interface {
	T() string
}

type Sub struct {
	S string
}

func (s *Sub) T() string {
	panic("implement me")
}

func TT() []Super {
	ts := make([]Super, 10)
	for i := 0; i < 10; i++ {
		ts[i] = &Sub{"fuck"}
	}
	return ts
}

// append相关操作
func T_Append() {
	ints := []int{1, 2, 3}
	ints2 := ints[0:1]
	println(len(ints2))
	//ints2 = append(ints2, 4, 5, 6)
	// 上面这种同时append多个值的方式会直接导致引用的array改变，所以ints不会改变
	ints2 = append(ints2, 4)
	ints2 = append(ints2, 5)
	ints2 = append(ints2, 6)
	println(util.ToJson(ints))
	println(util.ToJson(ints2))
}

// []*T指针相关操作
func T_Pointer() {
	type Person struct {
		Name string
	}
	ps := []*Person{{"fuck"}, {"happy"}, {"test"}}
	for _, p := range ps {
		p.Name = "haha"
	}
	for _, p := range ps {
		println(p.Name)
	}

	p := Person{"fuck"}
	p.Name = "hah"
	println(p.Name)
}

func main() {
	T_Append()
}
