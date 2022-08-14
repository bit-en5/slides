package main

import (
	"fmt"
	"unicode/utf8"
)

func loop() {
	//START OMIT
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	// While loop
	n := 0
	for n < 5 {
		n++
	}

	// Infinite loop
	for {
		break // allow to quit the loop
	}

	s := "how to iterate over 界 UTF8 string? 🙂"
	fmt.Printf("length: %d/ Number of chars: %d\n", len(s), utf8.RuneCountInString(s))
	for i, char := range s { // HL
		fmt.Printf("rune %d at position %d -> character %s\n", char, i, string(char))
	}
	//END OMIT
}
