package operate

import "os/exec"

func Status() (out []byte, err error) {
	open := exec.Command("./"+ScriptBard, "status")
	open.Dir = ScriptDir
	// 阻塞执行， 这个命令很快可以阻塞
	return open.Output()
}