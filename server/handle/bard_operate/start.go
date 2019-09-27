package bard_operate

import (
	"bard-gui/server/operate"
	"fmt"
	"net/http"
)

//@uri	/bard/start
func Start(w http.ResponseWriter, req *http.Request) {
	// TODO 开启bard
	rsp := Do(operate.Open)

	fmt.Fprint(w, rsp)
}

