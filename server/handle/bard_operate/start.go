package bard_operate

import (
	"fmt"
	"net/http"
)

func Start(w http.ResponseWriter, req *http.Request) {
	// TODO 开启bard

	fmt.Fprint(w, "open bard")
}

