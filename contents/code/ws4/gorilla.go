package main

//START OMIT
import (
	"net/http"

	"github.com/gorilla/mux" // HL
)

func main() {
	r := mux.NewRouter() // .StrictSlash(true)  // HL

	r.Handle("/", http.FileServer(http.Dir("./assets")))
	r.HandleFunc("/api/hello", sayHello)
	r.HandleFunc("/api/greeting", jsonResponse)
	r.HandleFunc("/api/info", handleRequest)
	r.HandleFunc("/api/deploy", deployment)

	routes := addMiddlewaresToRoutes(r)

	if err := http.ListenAndServe(":8080", routes); err != nil {
		panic(err)
	}
}

//END OMIT

func sayHello(w http.ResponseWriter, r *http.Request)      {}
func jsonResponse(w http.ResponseWriter, r *http.Request)  {}
func handleRequest(w http.ResponseWriter, r *http.Request) {}
func deployment(w http.ResponseWriter, r *http.Request)    {}
func addMiddlewaresToRoutes(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}
