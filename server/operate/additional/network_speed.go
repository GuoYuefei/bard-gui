package additional

import (
	"bard-gui/server/operate"
	"fmt"
	"os/exec"
	"runtime"
)

const bardConfig = "../../../bard/client/debug/config/config.yml"

// 利用ping网络测速

func NetworkSpend() (string, error) {
	// 先读取配置bard的配置文件
	config, err := operate.ParseConfig(bardConfig)
	if err != nil {
		return err.Error(), err
	}
	serverIP := config.GetServers()[0]
	fmt.Println(serverIP)
	cmd := []string{"ping"}

	if runtime.GOOS == "windows" {
		cmd = append(cmd, "-n 5")
	} else {
		cmd = append(cmd, "-c 5")
	}
	cmd = append(cmd, serverIP)

	ping := exec.Command(cmd[0], cmd[1:]...)

	output, err := ping.Output()
	if err != nil {
		return err.Error(), err
	}

	return string(output), nil
}
