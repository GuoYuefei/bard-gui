package bard_operate

import (
	"fmt"
	"net/http"
)

func Restart(w http.ResponseWriter, req *http.Request) {
	// TODO 重启bard

	fmt.Fprint(w, "restart bard")
}