package bard_operate

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// 修改bard的主配置文件
const BardConfig = "./bard/client/debug/config/config.yml"

//@uri /bard/config/update
func ConfigUpdate(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("contentType", "application/json;charset=utf-8")
	defer req.Body.Close()

	data, e := ioutil.ReadAll(req.Body)
	if e != nil {
		// can not read data from request
		fmt.Fprint(w, "{\"code\":1}")
		return
	}
	fmt.Println(string(data))
	e = ioutil.WriteFile(BardConfig, data, os.ModePerm)

	if e != nil {
		// can not write data to config file
		fmt.Fprint(w, "{\"code\":2}")
		return
	}

	fmt.Fprint(w, "{\"code\":0}")
}

//@uri /bard/config/get
func ConfigGet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadFile(BardConfig)
	if err != nil {
		fmt.Fprint(w, "error")
		return
	}
	fmt.Println(string(data))

	fmt.Fprint(w, string(data))
}
