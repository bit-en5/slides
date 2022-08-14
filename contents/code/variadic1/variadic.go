package main

import "fmt"

//START1 OMIT

func Sum(nums ...int) {
	fmt.Println(len(nums), "numbers received")

	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

//END1 OMIT
