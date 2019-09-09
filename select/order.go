package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	intChans := make(chan int, 1)
	stringChans := make(chan string, 1)
	intChans <- 1
	stringChans <- "hello"
	// 顺序随机
	select {
	case value := <-intChans:
		fmt.Println(value)
	case value := <-stringChans:
		panic(value)
	}
}
