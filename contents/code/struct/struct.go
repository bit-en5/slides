package main

import "fmt"

//START1 OMIT

type Person struct {
	Name string
	Age  int
}

//END1 OMIT

//START2 OMIT

func printPerson() {
	user := Person{
		Name: "David",
		Age:  42,
	}

	fmt.Println(user)
}

//END2 OMIT

//START3 OMIT

type Rectangle struct {
	Width  int
	Height int
}

//END3 OMIT
