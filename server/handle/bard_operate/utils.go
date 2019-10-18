package bard_operate

import (
	"bard-gui/server/operate"
	"fmt"
)

// 定义反馈信息
const (
	success = "{\"code\":0}"
	failed = "failed"
)

func Do(cmdFunc operate.CmdFunc) string {
	var rsp string
	out, err := cmdFunc()

	fmt.Println("out:", string(out), "err:", err)
	if err == nil {
		if ok, message := IsSuccess(out); ok {
			rsp = "{\"code\": 0, \"message\":\""+string(message)+"\"}"
		} else {
			rsp = "{\"code\": 1, \"message\":\""+string(message)+"\"}"
		}
	} else {
		fmt.Println(err)
		// 这个错误前端会捕捉
		rsp = err.Error()
	}
	return rsp
}


// return bool, message
func IsSuccess(out []byte) (bool, []byte) {
	message := out[3:]
	if out[1] == '0' {
		// [0]xxx 这种情况认为是成功执行脚本的了
		return true, message
	}
	return false, message
}
