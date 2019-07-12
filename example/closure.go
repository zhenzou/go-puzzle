package main

func Add() func(x int) int {
	n := 10

	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	adder := Add()
	println(adder(100))
	println(adder(100))
}
