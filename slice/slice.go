package main

import (
	"reflect"

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
func Append() {
	arr1 := []int{1, 2, 3}
	arr2 := arr1[0:1]
	println(len(arr2))
	// 这种append多个值的直接导致引用的array发生resize操作，后续修改的是新的array，所以ints不会改变
	arr2 = append(arr2, 4, 5, 6)

	println(reflect.DeepEqual(arr1, arr1))

	arr1 = []int{1, 2, 3}
	arr2 = arr1[0:1]
	arr2 = append(arr2, 4)
	arr2 = append(arr2, 5)
	arr2 = append(arr2, 6)

	println(reflect.DeepEqual(arr1, arr2))

	println(util.ToJson(arr1))
	println(util.ToJson(arr2))
}

type Person struct {
	Name string
}

// []*T指针相关操作
func TraversePtrSlice() {
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

func TraverseStructSlice() {
	ps := []Person{{"fuck"}, {"happy"}, {"test"}}
	for _, p := range ps {
		p.Name = "haha"
	}
	for _, p := range ps {
		//不会修改ps中元素的值，copy
		println(p.Name)
	}
}

func Make() {
	// 只有参数，同时声明了容量以及长度，会初始化为0
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	println(util.ToJson(s))
	s1 := make([]int, 0, 5)
	s1 = append(s1, 1, 2, 3)
	println(util.ToJson(s1))
}

func main() {
	Make()
	Append()
	TraversePtrSlice()
}
