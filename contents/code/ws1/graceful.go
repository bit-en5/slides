package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	*http.Server
}

//START OMIT

func (srv *Server) run() {
	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 10)                     // HL
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // HL
	<-quit                                               // HL

	// Maximum allowed time for the current requests to complete
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// This method stops taking new requests and finishes the currently handled requests
	err := srv.Shutdown(ctxShutDown)
	if err != nil {
		panic(err)
	}
}

//END OMIT
