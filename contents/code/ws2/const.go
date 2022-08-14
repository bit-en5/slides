package main

import "net/http"

const (
	port = ":8080"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./assets")))

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
