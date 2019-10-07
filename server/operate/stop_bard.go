package operate

import "os/exec"

func Stop() ([]byte, error) {
	cmd := append(ScriptBard, "stop")
	stop := exec.Command(cmd[0], cmd[1:]...)
	stop.Dir = ScriptDir
	// 阻塞执行， 这个命令很快可以阻塞
	return stop.Output()
}

