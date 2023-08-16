package main

func main() {
	ch := make(chan string)
	ch <- "test1"
	println("End")
}
