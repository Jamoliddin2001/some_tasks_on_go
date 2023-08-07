package main

import "fmt"

func main() {

	/*a := 129812324

	if a%2 == 0 {
		fmt.Print("Chetniy")
	} else {
		fmt.Println("NeChetniy")
	}*/

	//fmt.Print(10 >> 2)

	/*var array [5]int
	array[0] = 1
	//TODO
	fmt.Print(array[0])*/

	/*if 4 > 5 {
		fmt.Print("ok")
	} else {
		fmt.Println("no")
	}*/

	/*a := 10
	switch a {
	case 9:
		fmt.Println("a = 9")
		break
	case 8:
		fmt.Println("a = 8")
		break
	case 7:
		fmt.Println("a = 7")
		break
	default:
		fmt.Print("its not time")
	}*/

	/*var a = []int{}  // length = 0
	a = append(a, 5) // length = 1
	a = append(a, 6)
	a = append(a, 7)
	a = append(a, 8)
	a = append(a, 9)
	fmt.Print(a[4])*/

	/*var myArray = []int{-10, 5, 6, 7, 8, 2, 9, 2, 5, 4}

	for i := 0; i < len(myArray); i++ {
		if myArray[i]%2 == 0 {
			myArray[i] = myArray[i] / 2
		} else {
			myArray[i] = myArray[i] * myArray[i]
		}
	}

	for i := 0; i < len(myArray); i++ {
		fmt.Print(myArray[i], " ")
	}*/

	var myArray = []int{-10, 5, 6, 7, 8, 2, 9, 2, 5, 4}

	var s = 0
	for i := 0; i < len(myArray); i++ {
		if myArray[i]%2 == 1 {
			s += myArray[i]
		}
	}

	fmt.Print(s)
}
