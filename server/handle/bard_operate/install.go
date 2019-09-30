package bard_operate

import (
	"bard-gui/server/operate"
	"fmt"
	"net/http"
)

//@uri	/bard/install
func Install(w http.ResponseWriter, req *http.Request) {
	rsp := Do(operate.Install)
	fmt.Fprint(w, rsp)
}