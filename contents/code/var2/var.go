package main

import "fmt"

//START2 OMIT

var (
	a int
	b string
	c bool
)

//END2 OMIT

func Var() {
	//START3 OMIT
	var i int
	var j int = 15
	var k = 15 // inference
	l := 31    // inference
	s1 := `line 1
	line 2
	line 3` // inference

	s2 := "Hello\nworld!\t\n\"你好世界\""
	世界 := "hello"                    // 🤬
	var piValue, piName = 3.14, "pi" // inference 😐
	r := 'a'                         // 🛑 it's a rune, not a string
	//END3 OMIT

	fmt.Println(i, j, k, l, s1, s2, r, piValue, piName, 世界)
}
