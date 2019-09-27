package operate

import "os/exec"

func Open() (out []byte, err error) {
	open := exec.Command("./"+ScriptBard, "start")
	open.Dir = ScriptDir
	// 阻塞执行， 这个命令很快可以阻塞
	return open.Output()

}
