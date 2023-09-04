package main

/*
Описать функцию DigitCount(K) целого типа, находящую количество цифр целого положительного числа K.
Используя эту функцию, найти количество цифр для каждого из пяти данных целых положительных чисел.
*/
func digitCount(k int) int {
	s := 0
	for k != 0 {
		k = k / 10
		s++
	}
	return s
}

func main() {
	println("digitCount: ", digitCount(1726378123))
}
