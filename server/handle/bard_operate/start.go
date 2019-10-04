package bard_operate

import (
	"bard-gui/server/operate"
	"fmt"
	"net/http"
)

//@uri	/bard/start
func Start(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// TODO 系统判定
	rsp := Do(operate.Open)

	fmt.Fprint(w, rsp)
}

