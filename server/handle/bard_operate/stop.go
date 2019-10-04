package bard_operate

import (
	"bard-gui/server/operate"
	"fmt"
	"net/http"
)

//@uri	/bard/stop
func Stop(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	rsp := Do(operate.Stop)

	fmt.Fprint(w, rsp)
}