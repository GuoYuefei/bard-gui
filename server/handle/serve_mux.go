package handle

import (
	"fmt"
	"net/http"
)

var Mux = http.NewServeMux()

func init() {
	Mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprint(w, "WelCome to the home page!")
	})
	Mux.HandleFunc("/api", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "WelCome to the api page!")
	})
}