package bard_operate

import (
	"fmt"
	"net/http"
)

func Stop(w http.ResponseWriter, req *http.Request) {
	// TODO 关闭bard

	fmt.Fprint(w, "close bard")
}