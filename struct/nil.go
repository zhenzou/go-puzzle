package main

type Nil struct {
	Name string
}

func (n *Nil) NotPanic() {
	println("NotPanic")
}

func (n *Nil) Panic() {
	println("Panic:", n.Name)
}

func main() {
	var n *Nil

	n.NotPanic()
	n.Panic()
}
