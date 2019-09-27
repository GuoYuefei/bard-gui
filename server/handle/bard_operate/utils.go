package bard_operate

import (
	"bard-gui/server/operate"
	"fmt"
)

func Do(cmdFunc operate.CmdFunc) string {
	var rsp string
	out, err := cmdFunc()

	fmt.Println("out:", string(out), "err:", err)
	if err == nil {
		rsp = "success"
	} else {
		fmt.Println(err)
		rsp = err.Error()
	}
	return rsp
}
