package main

type Person struct {
	Name string
}

func (p Person) SetName1(name string) {
	p.Name = name
}

func (p *Person) SetName2(name string) {
	p.Name = name
}

func main() {
	p1 := &Person{"fuck1"}
	p1.SetName1("haha")
	println(p1.Name) // fuck1

	p2 := Person{"fuck2"}
	// 编译器会自动转换成(&p2).SetName2()这种形式，所以，p2必须是addressable的
	p2.SetName2("haha")
	println(p2.Name)
}
