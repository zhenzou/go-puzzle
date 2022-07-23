package main

type Person interface {
	Say()
}

type person1 struct{}

// Say 这种实现方式，ptr 类型也会实现 Person interfa
func (p person1) Say() {
	println("say 2")
}

type person2 struct{}

func (p *person2) Say() {
	println("say 1")
}

func main() {
	// person1
	Person(person1{}).Say()
	Person(&person1{}).Say()

	// peron2
	// Person(person2{}).Say()
	Person(&person2{}).Say()
}
