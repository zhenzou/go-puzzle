package main

import "fmt"

func stdInput() {
	var s string
	for {
		_, err := fmt.Scanf("%s", &s)
		if err != nil {
			break
		}
		fmt.Println("scan ", s)
	}
}

func main() {
	stdInput()
}
