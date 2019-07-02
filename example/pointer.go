package main

func f(point *int) {
	var data int
	point = &data
}

func f1(point *int) {
	var data int
	*point = data
}

func main() {
	i := 100

	f(&i)
	println(i)

	f1(&i)
	println(i)
}
