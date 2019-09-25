package main

import (
	"net/http"
	"server/handle"
	"time"
)

func main() {
	s := &http.Server{
		Addr:              ":2019",
		Handler:           handle.Mux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	s.ListenAndServe()
}
