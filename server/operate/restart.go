package operate

import "os/exec"

func Restart() (out []byte, err error) {
	cmd := append(ScriptBard, "restart")
	open := exec.Command(cmd[0], cmd[1:]...)
	open.Dir = ScriptDir
	// 阻塞执行， 这个命令很快可以阻塞
	return open.Output()
}