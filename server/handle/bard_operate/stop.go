package bard_operate

import (
	"bard-gui/server/operate"
	"fmt"
	"net/http"
)

func Stop(w http.ResponseWriter, req *http.Request) {
	rsp := Do(operate.Stop)

	fmt.Fprint(w, rsp)
}