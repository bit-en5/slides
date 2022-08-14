package main

import "net/http"

func main() {
	r := http.NewServeMux() // HL
	r.HandleFunc("/api/hello", sayHello)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
