package main

import "log"

// 验证defer调用的函数为nil的时候的表现
// deferFunc(3)=7
func deferFunc(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()
	// 在这行，f已经确定是nil，所以defer的参数已经确定了，一个nil的func
	// 原理：编译器调用，deferproc，把f以及相关参数 作为 defer的参数传入，而这时候的f=nil
	defer f()
	f = func() {
		println("fuck")
		r += 2
	}
	return n + 1
}

// defer 会在当前函数的最后一行后执行
func main() {
	func() {
		defer func() {
			log.Println("defer inner")
		}()
	}()

	defer log.Println("defer")

	log.Println("main")
}
