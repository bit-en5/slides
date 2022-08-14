package main

import (
	"fmt"
)

func main() {
	var x uint = 9  // 0000 1001
	var y uint = 65 // 0100 0001

	fmt.Println(x & y)  // 1
	fmt.Println(x | y)  // 73
	fmt.Println(x ^ y)  // 72
	fmt.Println(x << 1) // 18
	fmt.Println(x >> 1) // 4
}
