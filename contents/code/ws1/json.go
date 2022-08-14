package main

import (
	"encoding/json"
	"net/http"
)

//START OMIT

func jsonResponse(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string `json:"msg"`
		For     string `json:"receiver"`
		IsMale  bool   `json:"isMale,omitempty"`
		Value   int
	}{
		Message: "Hello",
		For:     "Coco",
		IsMale:  true,
		Value:   1,
	}

	resp, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(resp)
	}
}

//END OMIT
