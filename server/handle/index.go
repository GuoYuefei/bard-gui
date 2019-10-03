package handle

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	// 允许接受各ip/domain:port处请求
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	fmt.Fprint(w, "WelCome to the home page!")
}
