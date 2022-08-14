package main

import "fmt"

func main() {
	s := "hello"
	test(&s)
	fmt.Println("result: ", s)
}

// Bad example
func test(txt *string) {
	fmt.Println(*txt)
	*txt = "coucou"
	fmt.Println(*txt)
}
