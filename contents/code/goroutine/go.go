package main

import "fmt"

func Goroutine() {
	//START OMIT
	go fmt.Println("Hello")
	go myFunc()
	//END OMIT
}

func myFunc() {}
