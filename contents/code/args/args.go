package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Args[1:])

	for i, arg := range os.Args {
		if i == 0 {
			continue
		}

		fmt.Println(i, arg)
	}
}
