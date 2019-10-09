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
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      600 * time.Second,			// 安装花费时间大，根据网络情况而定
	}
}
