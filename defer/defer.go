package main

import "log"

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
