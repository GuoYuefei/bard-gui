package bard_operate

import (
	"fmt"
	"net/http"
)

// 修改bard的主配置文件

func ConfigUpdate(w http.ResponseWriter, req *http.Request) {
	// TODO 根据json，生成yml，更新配置文件


	fmt.Fprint(w, "update bard config")
}

func ConfigGet(w http.ResponseWriter, req *http.Request) {
	// TODO 根据yaml，生成json，返回json信息


	fmt.Fprint(w, "get bard config")
}
