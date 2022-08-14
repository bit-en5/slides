package main

//START2 OMIT

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

//END2 OMIT
