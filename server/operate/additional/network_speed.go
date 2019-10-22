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

type ICMP_Head struct {
	Type uint8
	Code uint8
	CheckSum []byte
	Identifier []byte
	SeqNum []byte
	Data []byte
}

func NewICMP_Head(t uint8, code uint8, ide uint16, seq uint16) *ICMP_Head {

	return &ICMP_Head{
		Type:       t,
		Code:       code,
		CheckSum:   []byte{0, 0},
		Identifier: []byte{byte(ide>>8), byte(ide)},
		SeqNum:     []byte{byte(seq>>8), byte(seq)},
		Data:       []byte{},
	}
}

// icmp echo
// type 8 code 0
// icmp reply
// type 0 code 0


// 正常加法 但是溢出一次自动+1一次
func (i *ICMP_Head) DoCheckSum() uint16 {
	var sum uint16 = 0
	Data := []byte{
		i.Type, i.Code, 0x00, 0x00,
		i.Identifier[0], i.Identifier[1],
		i.SeqNum[0], i.SeqNum[1],
	}
	Data = append(Data, i.Data...)
	if len(Data) % 2 != 0 {
		Data = append(Data, 0)
	}
	for j := 0; j < len(Data)/2; j++ {
		temp := sum
		sum += uint16(Data[2*j])<<8 + uint16(Data[2*j+1])
		if sum < temp {				// 不能== 因为可能存在0
			// 溢出了
			sum++
		}
		//fmt.Printf("%d:\t sum is %x\n", j, sum)
	}
	return sum
}



func GoPing() {

}
