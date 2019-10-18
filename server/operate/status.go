package operate

import "os/exec"

func Status() (out []byte, err error) {
	cmd := append(ScriptBard, "status")
	status := exec.Command(cmd[0], cmd[1:]...)
	status.Dir = ScriptDir
	// 阻塞执行， 这个命令很快可以阻塞
	return status.Output()
}