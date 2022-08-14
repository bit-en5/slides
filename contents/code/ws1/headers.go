package main

import "net/http"

//START OMIT

// SetHeaders sets the response with the default http headers (security)
func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	w.Header().Add("Expires", "0")
	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0")
	w.Header().Add("Pragma", "no-cache")
	w.Header().Add("X-Frame-Options", "SAMEORIGIN")
	w.Header().Add("X-Xss-Protection", "1; mode=block")
	w.Header().Add("X-Content-Type-Options", "nosniff")
}

//END OMIT
