package main

type Fucker interface {
	Fuck()
}

type FuckerImpl1 struct {
}

func (FuckerImpl1) Fuck() {
	println("FuckerImpl1")
}

type FuckerImpl2 struct {
}

func (*FuckerImpl2) Fuck() {
	println("FuckerImpl2")
}

func TestFucker(fucker Fucker) {
	fucker.Fuck()
}

func main() {
	var fucker1 Fucker = FuckerImpl1{}
	fucker1.Fuck()

	//注意，实现接口的方法的接受者类型是否是指针
	var fucker2 Fucker = &FuckerImpl2{}
	fucker2.Fuck()

	TestFucker(fucker1)

	fucker3 := FuckerImpl2{}

	fucker4 := &FuckerImpl1{}
	TestFucker(&fucker3)
	// 这个没有报错，意思是编译器会自动去引用，但是不会自动取引用
	TestFucker(fucker4)
}
