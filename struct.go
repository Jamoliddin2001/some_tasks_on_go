package main

import "lesson1/newPackage"

type person struct {
	name string
	age  int
}

func main() {
	var tom person = person{"Tom", 23}
	var tomPointer = &tom
	(*tomPointer).age = 22
	println((*tomPointer).name, tomPointer.age)

	println(newPackage.HelloWorld())
}
