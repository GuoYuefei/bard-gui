package operate

import "os/exec"

func Restart() (out []byte, err error) {
	open := exec.Command("./"+ScriptBard, "restart")
	open.Dir = ScriptDir
	// 阻塞执行， 这个命令很快可以阻塞
	return open.Output()
}