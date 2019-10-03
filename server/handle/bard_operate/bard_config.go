package bard_operate

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 修改bard的主配置文件
const BardConfig = "./bard/client/debug/config/config.yml"

//@uri /bard/config/update
func ConfigUpdate(w http.ResponseWriter, req *http.Request) {
	// TODO 根据json，生成yml，更新配置文件


	fmt.Fprint(w, "update bard config")
}

//@uri /bard/config/get
func ConfigGet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadFile(BardConfig)
	if err != nil {
		fmt.Fprint(w, "error")
		return
	}

	fmt.Fprint(w, string(data))
}
