package operate

import (
	"os/exec"
	"runtime"
)

const (
	ScriptDir = "./bard"
	ScriptBard_uix = "./bard.sh"
	ScriptInstall = "./install-bard.bash"
	ScriptBardWin = "bard.cmd"
)

var ScriptBard []string

func init() {
	ScriptBard = append(ScriptBard, ScriptBard_uix)

	if runtime.GOOS == "windows" {
		wincmd := make([]string, 2)
		wincmd[0] = "cmd"
		wincmd[1] = "/c"
		ScriptBard = append(wincmd, ScriptBardWin)
	} else {
		// 赋予执行权限
		chmod := exec.Command("chmod", "+x", ScriptBard[0], ScriptInstall)
		chmod.Dir = ScriptDir
		_ = chmod.Run()

	}
}

type CmdFunc = func() ([]byte, error)
