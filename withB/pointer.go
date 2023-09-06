package main

//var y int = 10

func copyPar(p *int) {
	*p++
	println("adress p=", p)
	println("In function = ", *p)
}

func main() {
	x := 10
	copyPar(&x)
	println("adress x=", &x)
	println("After function =", x)
	//var p *int
	//p = &x
	//println("p = ", p)
	//println("&x = ", &x)
	//println("x= ", x)
	//println("*p = ", *p)
	//x++
	//println("-----------------")
	//println("p = ", p)
	//println("&x = ", &x)
	//println("x = ", x)
	//println("*p = ", *p)
	//*p++
	//println("-----------------")
	//println("p = ", p)
	//println("&x = ", &x)
	//println("x = ", x)
	//println("*p = ", *p)
	//copyPar(x)
	//println(x)
}
