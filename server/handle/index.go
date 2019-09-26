package handle

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	fmt.Fprint(w, "WelCome to the home page!")
}
