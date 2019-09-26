package main

import (
	. "bard-gui/server"
	"fmt"
	"log"
	"strings"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	//S.Addr="127.0.0.1:2048"			// 地址可变
	panic(ListenAndOpenBrowser())
}

func ListenAndOpenBrowser() error {
	var url = "http://"
	if strings.Index(S.Addr, ":") == 0 {
		url = url + "localhost" + S.Addr
	} else {
		url += S.Addr
	}

	if err := openUrl(url); err != nil {
		log.Println("打开游览器发生错误:", err)
		fmt.Println("请手动打开游览器输入：", url)
	}

	// 开启服务			服务不大，所以先开游览器再开服务器
	return S.ListenAndServe()
}
