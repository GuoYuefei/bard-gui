package bard_operate

import (
	"bard-gui/server/operate"
	"fmt"
	"net/http"
)

//@uri	/bard/install
func Install(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	rsp := Do(operate.Install)
	fmt.Println(rsp)
	fmt.Fprint(w, rsp)
}