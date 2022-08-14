package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

//START OMIT

func deployment(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // HL

	var msg string
	switch time.Now().Weekday() {
	case time.Friday:
		msg = "NEVER deploy on Friday"
	case time.Saturday, time.Sunday:
		msg = "Avoid to deploy today"
	default:
		msg = "Happy deployment 🙂"
	}

	data := struct{ For, Message string }{For: strings.TrimSpace("Hello " + name), Message: msg}

	resp, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(resp)
	}
}

//END OMIT
