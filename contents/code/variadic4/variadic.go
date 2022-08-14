package main

//START4 OMIT

func Sum(nums ...int) (int, int) {
	nb := len(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	return nb, total
}

//END4 OMIT
