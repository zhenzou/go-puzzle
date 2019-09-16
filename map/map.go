package main

import "fmt"

func pass(m map[string]string) map[string]string {
	m = map[string]string{"name": "test2"}
	return m
}

func clear(m map[string]string) {
	for k, _ := range m {
		delete(m, k)
	}
}

func main() {
	m1 := map[string]string{"name": "test"}
	m2 := pass(m1)
	m3 := make(map[string]string)
	m3["name"] = "test3"
	m4 := pass(m3)

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(m4)

	m3["sex"] = "female"
	clear(m3)
}
