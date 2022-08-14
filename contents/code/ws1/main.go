package main

import "net/http"

func main() {
	//START2 OMIT
	http.Handle("/", http.FileServer(http.Dir("./assets")))

	//START1 OMIT
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	//END OMIT
}
