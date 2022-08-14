package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

//START1 OMIT

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/api/hello", sayHello)

	routes := addMiddlewaresToRoutes(r) // HL

	if err := http.ListenAndServe(":8080", routes); err != nil {
		panic(err)
	}
}

//END1 OMIT
//START2 OMIT

func addMiddlewaresToRoutes(h http.Handler) http.Handler {
	// securityHandler adds some headers to the response
	h = securityHandler(h)

	// CORSHandler allows some domains to use the application
	h = CORSHandler(h)

	// RecoverPanic avoid the application to crash in case of panic
	h = recoverPanic(h)

	return h
}

//END2 OMIT
//START3 OMIT

// timeOutHandler adds a max time to handle the request before timeout
func timeOutHandler(next http.Handler, timeout int) http.Handler {
	f := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Consider timeout only if the duration is bigger than 0
		if timeout > 0 {
			ctx := context.Background()

			ctxTimeOut, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
			defer cancel()

			r = r.WithContext(ctxTimeOut)
		}

		// Let's work
		next.ServeHTTP(w, r)
	})

	return f
}

//END3 OMIT
//START4 OMIT

func securityHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		w.Header().Add("Expires", "0")
		w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0")
		w.Header().Add("Pragma", "no-cache")
		w.Header().Add("X-Frame-Options", "SAMEORIGIN")
		w.Header().Add("X-Xss-Protection", "1; mode=block")
		w.Header().Add("X-Content-Type-Options", "nosniff")

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

//END4 OMIT
//START5 OMIT

// CORSHandler middleware checks each request for it's origin and sets Access-Control headers
func CORSHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", "localhost:*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, enctype")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true") // Allow cookie to be sent
		}

		// Stop here if its Preflighted OPTIONS request
		if r.Method == "OPTIONS" {
			return
		}

		// Lets work
		next.ServeHTTP(w, r)
	})
}

//END5 OMIT
//START6 OMIT

// recoverPanic allows to recover in case of panic
func recoverPanic(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Print(fmt.Errorf("panic: %+v", r))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

//END6 OMIT

func sayHello(w http.ResponseWriter, r *http.Request) {}
