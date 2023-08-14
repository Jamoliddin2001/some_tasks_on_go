package main

import "fmt"

/*
	func name_of_func(param1 type, ...) (returnType) {
		// TODO
	}
*/

func main() {
	fmt.Println(add(4, 5))
	fmt.Println(returnTwoParams(14, 15, "First", "Last"))
	fmt.Println(multiply(4, 7))
	fmt.Println(factorial(5))
}

func add(x, y int) int {
	return x + y
}

func returnTwoParams(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}

func multiply(x, y int) (z int) {
	z = x * y
	return
}

// recursion
func factorial(n uint) uint {

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
