package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		go println(i, "!=", factorial1(i))
	}
	_, err := fmt.Scanln()
	if err != nil {
		return
	}
	fmt.Println("The End")
}

func factorial1(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial1(n-1)
}
