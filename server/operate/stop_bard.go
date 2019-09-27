package operate

import "os/exec"

func Stop() ([]byte, error) {
	stop := exec.Command("./"+ScriptBard, "stop")
	stop.Dir = ScriptDir
	// 阻塞执行， 这个命令很快可以阻塞
	return stop.Output()
}

