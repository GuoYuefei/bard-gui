package additional

import (
	"bard-gui/server/operate"
	"fmt"
	"net"
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

type ICMP struct {
	*net.IPConn
	Type uint8
	Code uint8
	CheckSum []byte
	Identifier []byte
	SeqNum []byte
	Data []byte
}

func NewICMP(t uint8, code uint8, ide uint16, seq uint16) *ICMP {

	return &ICMP{
		IPConn:		nil,
		Type:       t,
		Code:       code,
		CheckSum:   []byte{0, 0},
		Identifier: []byte{byte(ide>>8), byte(ide)},
		SeqNum:     []byte{byte(seq>>8), byte(seq)},
		Data:       []byte{},
	}
}

// 从接受到的bytes中生成ICMP
func NewICMPFromBytes(src []byte) *ICMP {
	if len(src) < 4 {
		return nil
	}
	var bs []byte
	if len(src) < 8 {
		bs = make([]byte, 8)
		copy(bs, src)
	} else {
		bs = src
	}

	return &ICMP{
		IPConn:     nil,
		Type:       bs[0],
		Code:       bs[1],
		CheckSum:   bs[2:4],
		Identifier: bs[4:6],
		SeqNum:     bs[6:8],
		Data:       bs[8:],
	}
}

func (icmp *ICMP)SetIPConn(conn *net.IPConn) {
	icmp.IPConn = conn
}

func (icmp *ICMP)SetIPConnByAddr(netProto string, laddr *net.IPAddr) error {
	conn, err := net.ListenIP(netProto, laddr)
	if err != nil {
		return err
	}
	icmp.SetIPConn(conn)
	return nil
}

func (icmp *ICMP)Bytes() []byte {
	bytes := []byte{icmp.Type, icmp.Code}
	bytes = append(bytes, icmp.CheckSum...)
	bytes = append(bytes, icmp.Identifier...)
	bytes = append(bytes, icmp.SeqNum...)
	bytes = append(bytes, icmp.Data...)

	return bytes
}

// icmp echo
// type 8 code 0
// icmp reply
// type 0 code 0


// 正常加法 但是溢出一次自动+1一次
func (icmp *ICMP) DoCheckSum() uint16 {
	var sum uint16 = 0
	Data := []byte{
		icmp.Type, icmp.Code, 0x00, 0x00,
		icmp.Identifier[0], icmp.Identifier[1],
		icmp.SeqNum[0], icmp.SeqNum[1],
	}
	Data = append(Data, icmp.Data...)
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

// ICMP应该可以自己发出请求吧
func (icmp *ICMP) Send() (int, error) {
	return icmp.Write(icmp.Bytes())
}


// receive messages
// 可以参考IPConn自己解析
//func (icmp *ICMP) Receive() *ICMP {
//	result := make([]byte, 64 * 1024)
//	l, addr, e := icmp.ReadFromIP(result)
//	if e != nil {
//		return nil
//	}
//
//}


func GoPing() {

}
