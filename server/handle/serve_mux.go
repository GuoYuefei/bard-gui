package handle

import (
	"bard-gui/server/handle/bard_operate"
	"net/http"
)

var Mux = http.NewServeMux()

func init() {
	// handle 配置
	Mux.HandleFunc("/", Home)
	Mux.HandleFunc("/bard/start", bard_operate.Start)
	Mux.HandleFunc("/bard/stop", bard_operate.Stop)
	Mux.HandleFunc("/bard/restart", bard_operate.Restart)
	Mux.HandleFunc("/bard/install", bard_operate.Install)
	Mux.HandleFunc("/bard/config/update", bard_operate.ConfigUpdate)
	Mux.HandleFunc("/bard/config/get", bard_operate.ConfigGet)
}
