package main

func If() {
	//START1 OMIT

	var (
		a, b int
		ok   bool
	)

	if a != b { // Do stuff 🙂
	}

	if ok { // Do stuff 🙂
	}

	if a == b { // 😐
		// Do stuff
	} else {
		// Do stuff
	}

	//END1 OMIT
}

func IfElse(name string) {
	//START2 OMIT

	// 🥴
	if len(name) == 1 {
		// Do stuff
	} else if len(name) == 2 {
		// Do stuff
	} else if len(name) == 3 {
		// Do stuff
	} else if len(name) == 4 {
		// Do stuff
	}

	//END2 OMIT
}
