package other

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
	获取bard插件的id列表
 */

const plugins = "./server/plugins/plugins.yml"

// @uri /plugins/get
func PluginsGet(w http.ResponseWriter, req *http.Request) {
	// 允许接受各ip/domain:port处请求
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, e := ioutil.ReadFile(plugins)
	if e != nil {
		fmt.Fprint(w, "error")
		return
	}

	fmt.Fprint(w, string(data))
}