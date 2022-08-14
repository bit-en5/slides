package main

import (
	"fmt"
)

//START1 OMIT

func Test1() {
	fmt.Println("hello")
}

//END1 OMIT
//START2 OMIT

func test2(name string) {
	fmt.Println("hello", name, "!")
}

//END2 OMIT
//START3 OMIT

func Test3(greeting string, firstname string, lastname string) {
	fmt.Println(greeting, firstname, lastname, "!")
}

//END3 OMIT

func anonymous() {
	//START5 OMIT

	func() {
		fmt.Println("Welcome!")
	}()

	//END5 OMIT

	//START6 OMIT

	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")

	//END6 OMIT
	//START7 OMIT

	value := func(n string) {
		fmt.Println("Welcome", n)
	}

	value("Gophers")
	value("Apprentice")
}

//END7 OMIT
