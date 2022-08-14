package main

import "net/http"

//START OMIT

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

//END OMIT
