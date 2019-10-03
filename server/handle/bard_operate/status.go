package bard_operate

import (
	"bard-gui/server/operate"
	"fmt"
	"net/http"
)

// 检查bard程序是否在运行

//@uri	/bard/status
func Status(w http.ResponseWriter, req *http.Request) {
	rsp := Do(operate.Stop)

	fmt.Fprint(w, rsp)
}
