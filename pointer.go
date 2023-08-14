package main

func main() {
	var x = 5
	var pointer *int
	pointer = &x
	println(pointer)

	println(increment(&x))
	println(x)
}

func increment(x *int) int {
	*x += 10
	return *x
}
