package bard_operate

import (
	"bard-gui/server/operate"
	"fmt"
	"net/http"
)

//@uri	/bard/restart
func Restart(w http.ResponseWriter, req *http.Request) {
	rsp := Do(operate.Restart)
	fmt.Fprint(w, rsp)
}