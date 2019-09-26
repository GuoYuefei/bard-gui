package main

import (
	"os/exec"
	"runtime"
)

// 开启软件时自动打开游览器的操作

var commands = map[string]string {
	"windows": "cmd /c start", 			// windows
	"darwin": "open",					// mac
	"linux": "xdg-open",				// linux
}

func openUrl(url string) error {
	openBrowser := exec.Command(commands[runtime.GOOS], url)
	return openBrowser.Run()
}