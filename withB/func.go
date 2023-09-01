package main

func hello() {
	println("Hello from func hello")
}

func square(x int) int {
	return x * x
}

func test(a string) {
	//TODO
	println("Text = ", a)
}

func sortArray(x []int) {
	//todo
}

func getNameAndAgeAfterOneYear(age int, name string) (int, string) {
	return age + 1, "Imran"
}

func main() {
	println("Hello akai Dilshod!")
	hello()
	var x = 234
	println(x, "^ 2 = ", square(x))
	test("Hello world!")
	var z, y = getNameAndAgeAfterOneYear(31, "Dilshod")
	println("Age after one year = ", z, " Name =", y)
}
