package main

// loop 1547447
//		1559817

func rec(x int) {
	println(x)
	if x < 100 {
		rec(x + 1)
	}
}
func factorial(x uint) uint {
	if x == 0 {
		return 1
	} else {
		return x * factorial(x-1)
	}
}

func main() {
	println(factorial(6))
}
