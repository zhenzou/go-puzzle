package main

func makeIntChan() <-chan int {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func recv() {
	intChan := makeIntChan()
	// 不能close
	for i := range intChan {
		println(i)
	}
}

func main() {
	recv()
}
