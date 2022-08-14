package main

import "time"

func Switch() {
	//START OMIT

	switch hour := time.Now().Hour(); {
	case hour < 6:
		// do stuff

	case hour < 12:
		// do stuff

	case hour > 18:
		// do stuff

	default:
		// more stuff
	}

	//END OMIT
}
