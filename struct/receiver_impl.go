package main

type Runner interface {
	Run()
}

type StructRunnerImpl struct {
}

func (StructRunnerImpl) Run() {
	println("StructRunnerImpl")
}

func (*StructRunnerImpl) Run2() {
	println("StructRunnerImplRun2")
}

type PtrRunnerImpl struct {
}

func (PtrRunnerImpl) Run2() {
	println("PtrRunnerImplRun2")
}

func (*PtrRunnerImpl) Run() {
	println("PtrRunnerImpl")
}

func TestRunner(fucker Runner) {
	fucker.Run()
}

// 结论：
// 1. 如果receiver是结构体：则 结构体和指针类型都实现了interface
// 2. 如果receiver是指针：则只有指针类型实现了interface
func main() {
	var structRunner1 Runner = StructRunnerImpl{}
	structRunner1.Run()

	structRunner2 := &StructRunnerImpl{}
	structRunner2.Run2()
	// 这个没有报错，编译器会自动解引用，但是不会自动取引用
	TestRunner(structRunner2)
	TestRunner(*structRunner2)

	//注意，实现接口的方法的接受者类型是否是指针
	var ptrRunner1 Runner = &PtrRunnerImpl{}
	ptrRunner1.Run()

	TestRunner(structRunner1)

	ptrRunner2 := PtrRunnerImpl{}
	// 这里必须取指针，不然编译报错
	TestRunner(&ptrRunner2)
	ptrRunner2.Run2()

	ptrRunner3 := &PtrRunnerImpl{}

	TestRunner(ptrRunner3)
	// 编译报错
	//TestRunner(*ptrRunner3)

}
