package bard_operate

import (
	"bard-gui/server/operate"
	"fmt"
)

// 定义反馈信息
const (
	success = "success"
	failed = "failed"
)

func Do(cmdFunc operate.CmdFunc) string {
	var rsp string
	out, err := cmdFunc()

	fmt.Println("out:", string(out), "err:", err)
	if err == nil {
		if IsSuccess(out) {
			rsp = success
		} else {
			rsp = failed
		}
	} else {
		fmt.Println(err)
		rsp = err.Error()
	}
	return rsp
}

func IsSuccess(out []byte) bool {
	if out[1] == '0' {
		// [0]xxx 这种情况认为是成功执行脚本的了
		return true
	}
	return false
}