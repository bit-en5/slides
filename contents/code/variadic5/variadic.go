package main

//START5 OMIT

func Sum(nums ...int) (nb, total int) {
	nb = len(nums)
	for _, num := range nums {
		total += num
	}
	return
}

//END5 OMIT
