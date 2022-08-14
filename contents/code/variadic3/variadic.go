package main

//START3 OMIT

func Sum(nums ...int) (total int) {
	for _, num := range nums {
		total += num
	}

	return
}

//END3 OMIT
