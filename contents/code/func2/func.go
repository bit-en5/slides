package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

//START4 OMIT

func Test3(greeting, firstname, lastname string) {
	fmt.Println(greeting, firstname, lastname, "!")
}

//END4 OMIT
//START8 OMIT

func main() {
	execute(doSomething, 123)
}

func doSomething(i int) {
	fmt.Printf("value is %v\n", i)
}

func execute(f func(int), val int) {
	f(val)
}

//END8 OMIT
//START9 OMIT

func Average(ages ...int) (av float64, err error) {
	nb := len(ages)
	if nb == 0 {
		return 0, errors.New("at least 1 age is required")
	}

	total := 0

	for _, age := range ages {
		if age == 0 {
			return 0, errors.New("age 0 is not allowed")
		}
		total += age
	}

	av = float64(total) / float64(nb)
	return
}

// END9 OMIT

func average() {
	// START10 OMIT
	average, err := Average(37, 25, 0)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	// END10 OMIT
	fmt.Println(average)
}
