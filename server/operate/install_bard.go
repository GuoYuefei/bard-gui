package operate

import "os/exec"

func Install() (out []byte, err error) {
	// 无论有没有在运行程序，先关闭在安装或重新安装
	Stop()
	cmd := append(ScriptBard, "install")
	open := exec.Command(cmd[0], cmd[1:]...)
	open.Dir = ScriptDir
	// 阻塞执行， 这个命令很快可以阻塞
	return open.Output()
}
