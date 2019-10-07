package bard_operate

import (
	"bard-gui/server/operate"
	"net/http"
)

//@uri	/bard/restart
func Restart(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	Do(operate.Stop)

	Start(w, req)
	//fmt.Fprint(w, rsp)
}