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

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live1() People {
	var stu *Student
	return stu
}

func live2() *Student {
	var stu *Student
	return stu
}

func main() {
	var n *Nil

	if live1() == nil {
		println("AAAAAAA")
	} else {
		println("BBBBBBB")
	}

	if live2() == nil {
		println("AAAAAAA")
	} else {
		println("BBBBBBB")
	}

	n.NotPanic()
	n.Panic()
}
