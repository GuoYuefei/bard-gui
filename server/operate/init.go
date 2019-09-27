package operate

import "C"
import "os/exec"

func init() {
	chmod := exec.Command("chmod", "+x", "./"+ScriptBard, "./"+ScriptInstall)
	chmod.Dir = ScriptDir
	_ = chmod.Run()
}

type CmdFunc = func() ([]byte, error)
