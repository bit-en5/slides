package main

import (
	"fmt"
	"net/http"
)

//START OMIT

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("User Agent:", r.UserAgent())
	fmt.Println("IP Address:", r.RemoteAddr)

	fmt.Printf("Request Method: %s\n", r.Method)
	fmt.Printf("Request RequestURI: %s\n", r.URL.RequestURI())
	fmt.Printf("Request URL path: %s\n", r.URL.Path)
	fmt.Printf("Request URL raw query: %s\n", r.URL.RawQuery)
}

//END OMIT
