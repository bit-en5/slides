package main

import "fmt"

//START1 OMIT

func main() {
	s := "hello"
	change(s)
	fmt.Println("final result: ", s)
}

func change(txt string) {
	fmt.Println(txt)
	txt = "coucou"
	fmt.Println(txt)
}

//END1 OMIT
