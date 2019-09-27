package main

import (
	"os/exec"
	"runtime"
)

// 开启软件时自动打开游览器的操作

var commands = map[string][]string {
	"windows": {"cmd", "/c", "start"}, 			// windows
	"darwin": {"open"},					// mac
	"linux": {"xdg-open"},				// linux
}

func openUrl(url string) error {
	command := commands[runtime.GOOS]
	command = append(command, url)
	// 0元素是命令，其余为参数 所有系统下命令至少有一个url的参数，所以一下切片写法合法
	openBrowser := exec.Command(command[0], command[1:]...)
	return openBrowser.Run()
}