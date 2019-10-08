package bard_operate

import (
	"bard-gui/server/operate"
	"fmt"
	"net/http"
)

// 检查bard程序是否在运行

//@uri	/bard/status
func Status(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	rsp := Do(operate.Status)

	fmt.Fprint(w, rsp)
}
