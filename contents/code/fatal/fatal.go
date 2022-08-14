package main

import (
	"log"
	"os"
)

func fatal() {
	//START OMIT
	log.Println("an error occured")
	os.Exit(1)
	//END OMIT
}
