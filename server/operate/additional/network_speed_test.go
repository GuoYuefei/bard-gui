package additional

import (
	"fmt"
	"testing"
)

func TestNetworkSpend(t *testing.T) {
	t.Log(NetworkSpend())

}

func TestICMP_Head_DoCheckSum(t *testing.T) {
	icmp := NewICMP(8, 0, 1, 1)
	icmp.Data = []byte{

		// type=0 code=0, ide=0, seq=0,     // will result should be 4e19, The inverse code is B1E6
		//0x45, 0x00, 0x00, 0x3c,
		//0x1c, 0x46,
		//0x40, 0x00,
		//0x40, 0x06, 0x00, 0x00,
		//0xac, 0x10, 0x0a, 0x63,
		//0xac, 0x10, 0x0a, 0x0c,

		// type=8, code=0, ide=1, seq=1,     // will result should be b2a5, The inverse code is 4d5a
		0X61, 0x62, 0X63, 0x64,
		0X65, 0x66, 0X67, 0x68,
		0X69, 0x6a, 0X6b, 0x6c,
		0X6d, 0x6e, 0X6f, 0x70,
		0X71, 0x72, 0X73, 0x74,
		0X75, 0x76, 0X77, 0x61,
		0X62, 0x63, 0X64, 0x65,
		0X66, 0x67, 0X68, 0x69,
	}
	sum := icmp.DoCheckSum()
	fmt.Printf("sum 16进制：%x,\nsum 10进制: %d,\nsum 反码16进制: %x\n", sum, sum, ^sum)
}
