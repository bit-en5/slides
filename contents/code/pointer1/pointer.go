package main

import "fmt"

func main() {
	//START1 OMIT

	a := 10
	fmt.Println("a =", a)
	fmt.Println("value of variable a is located at address", &a)

	//END1 OMIT
}

func getInfo() {
	//START2 OMIT

	a := 4
	info(&a)

	//END2 OMIT
}

//START3 OMIT

func info(x *int) {
	fmt.Println("value of variable is located at address", x)
}

//END3 OMIT
//START4 OMIT

func change(x *int) {
	*x = *x + 15
}

//END4 OMIT
