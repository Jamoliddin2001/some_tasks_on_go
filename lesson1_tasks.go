package main

func main() {

	/*for index, value := range array {
		fmt.Println(index, " -> ", value)
	}*/

	// #1 Дан массив размера N и целые числа K и L (1<K≤L≤N).
	// Найти сумму всех элементов массива, кроме элементов с номерами от K до L включительно.

	/*var array = []int{1, 4, 3, 432, 6, 6, 234, 2, -3, 4, -4, 2, -2}
	var k = 5
	var l = 7
	s := 0
	for i := 0; i < len(array); i++ {
		if i < k-1 || i > l-1 {
			s += array[i]
		}
	}
	fmt.Print("S = ", s)*/

	// #2 Дан массив ненулевых целых чисел размера N. Проверить, чередуются ли в нем положительные и отрицательные числа.
	// Если чередуются, то вывести 0, если нет, то вывести порядковый номер первого элемента, нарушающего закономерность.

	/*var array = []int{1, -4, 3, -432, -6, -6, 234, -2, 3, -4, 4, -2, 2}
	for i := 0; i < len(array)-1; i++ {
		fmt.Println("i = ", i, " array[", i, "] = ", array[i], " array[", i+1, "] = ", array[i+1])
		if array[i]*array[i+1] > 0 {
			fmt.Println("if check")
			fmt.Println("i = ", i+1)
			return
		}
	}
	fmt.Println(0)*/

	// #3 Дан целочисленный массив размера N, все элементы которого упорядочены (по возрастанию или по убыванию).
	// Найти количество различных элементов в данном массиве.

	/*var array = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	k := 0
	for i := 0; i < len(array)-1; i++ {
		if array[i] != array[i+1] {
			k++
		}
	}
	k++
	fmt.Println(k)*/

	// #4 Дан массив размера N, все элементы которого, кроме последнего, упорядочены по возрастанию.
	// Сделать массив упорядоченным, переместив последний элемент на новую позицию.

	/*var array = []int{1, 2, 4, 6, 6, 6, 7, 8, 11, 12, 12, 14, 17} // 3, 4, 6, 6, 6, 7, 8, 11, 12, 12, 14
	lastElement := array[len(array)-1]
	pos := 0
	for pos < len(array)-1 {
		if lastElement <= array[pos] {
			break
		}
		pos++
	}
	for i := len(array) - 1; i > pos; i-- {
		array[i] = array[i-1]
	}
	array[pos] = lastElement
	fmt.Println(array)*/

	/*var myArray = []int{1, 3, 4, 6, 1, 3, 2, 5, 7, -2, 6}
	for i := 0; i < len(myArray)-1; i++ {
		for j := i + 1; j < len(myArray); j++ {
			if myArray[i] > myArray[j] {
				t := myArray[i]
				myArray[i] = myArray[j]
				myArray[j] = t
			}
		}
	}

	fmt.Println(myArray)*/

	// refresh

}
