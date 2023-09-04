package main

import "fmt"

func main() {

	f := func(x, y int) int { return x + y }
	fmt.Println(f(3, 4))
	fmt.Println(f(6, 7))

}
