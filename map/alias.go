package main

import (
	"net/http"
)

func main() {

	// still need init
	var header http.Header

	header.Add("test", "test")

	println(header.Get("test"))
}
