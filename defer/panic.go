package main

import (
	"errors"
	"reflect"
)

func init() {
	println("init")
	defer println("init over")
}

func changeResult() (result int) {
	defer func() {
		result++
	}()
	return result
}

func changeResult2() (r int) {
	t := 5
	defer func() {
		r = t + 5
	}()
	return r //10 ,zai defer中可以修改r
}

func explain() (r int) {
	t := 5
	r = t //赋值指令
	{ //defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
		t = t + 5
	}
	return //空的return指令
}

func explain2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t //5
}

func explain3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func deferf() {
	i := 10
	defer func(i int) {
		println(2)
		println(i)
	}(i + 2)
	defer func(i int) {
		println(1)
		println(i)
	}(i + 2)

	i += 10

	println(i)
}

// panic在defer后执行
func panicdefer() {
	defer func() { println("打印前") }()
	defer func() { println("打印中") }()
	defer func() { println("打印后") }()

	panic("触发异常")
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func TestDeferChangeResult() {
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}

func DeferPanic() {
	defer func() {
		// recover的结果是panic的参数
		if err := recover(); err != nil {
			println(err)
			println(reflect.TypeOf(err).Name())
		} else {
			println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("panic")
}

// 只有最后一个panic会传回来，所以在defer中的panic会传回来
func TestDeferPanic() {
	defer func() {
		if err := recover(); err != nil {
			println("++++")
			f := err.(func() string)
			println(f())
			println(err, f(), reflect.TypeOf(err).Kind().String())
		} else {
			println("fatal")
		}
	}()

	defer func() {
		panic(func() string {
			return "defer panic"
		})
	}()
	panic("panic")
}

func DeferFunc() {
	defer func() {
		if err := recover(); err != nil {
			println(err)
		}
	}()
	panic(errors.New("test"))
}

func main() {

	//println(f1())
	//println(f2())
	//println(explain3())
	//println("fuck")
	panicdefer()
	//DeferPanic()
	//TestDeferPanic()
	//DeferFunc()
}
