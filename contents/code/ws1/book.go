package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//START OMIT

func getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isbn := vars["isbn"]

	w.Header().Set("Content-Type", "plain/text")
	w.Write([]byte("you want the book with isbn: " + isbn))
}

//END OMIT
