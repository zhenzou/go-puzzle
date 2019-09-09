package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func _copy() {
	a := 1
	b := 2
	// 函数的参数，即时计算，即时复制过去
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

/*
10 1 2 3
20 0 2 2
2  0 2 2
1  1 3 4
*/

func calc1(index string, a, b *int) *int {
	ret := *a + *b
	fmt.Println(index, *a, *b, ret)
	return &ret
}

func _ref() {
	a := 1
	b := 2
	// 函数的参数，即时计算，值传递
	defer calc1("1", &a, calc1("10", &a, &b))
	a = 0
	defer calc1("2", &a, calc1("20", &a, &b))
	b = 1
}

/*
10 1 2 3
20 0 2 2
2  0 2 2
1  0 3 3
*/

func main() {
	_ref()
}
