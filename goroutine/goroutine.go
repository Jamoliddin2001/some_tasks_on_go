package main

func main() {

	println("start")
	ch := make(chan int)
	go test(ch)
	println(<-ch)
	println("end")
}

func test(ch chan int) {
	println("goroutine start")
	ch <- 5
	for i := 0; ; i++ {
		println("i = ", i, "end")
	}
}
