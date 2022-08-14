package main

import "fmt"

//START1 OMIT

// Encapsulation with closure
func New() (Count func()) {
	n := 0
	return func() {
		n++
		fmt.Println(n)
	}
}

func main() {
	f1 := New()
	f1() // 1
	f1() // 2
	f1() // 3
}

//END1 OMIT
