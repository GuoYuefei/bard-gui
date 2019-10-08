package server

import (
	"bard-gui/server/handle"
	"net/http"
	"time"
)

var S *http.Server

func init() {
	S = &http.Server{
		Addr:              ":2019",
		Handler:           handle.Mux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      120 * time.Second,
	}
}
