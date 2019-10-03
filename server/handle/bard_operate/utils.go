package bard_operate

import (
	"bard-gui/server/operate"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
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


// 从bard项目copy过来的 管理yaml配置文件

type User struct {
	Username string						`yaml:"username"`
	Password string						`yaml:"password"`
	ComConfig *CommConfig					`yaml:"com_config"`

	// 以上基础信息
}

type CommConfig struct {
	// 插件id信息
	Plugins []string					`yaml:"plugins"`
	// 传输控制子协议id
	TCSP string							`yaml:"TCSP"`
}


type Config struct {
	Server interface{} 		`yaml:"server"`
	ServerPort int			`yaml:"server_port"`
	LocalPort int 			`yaml:"local_port"`
	LocalAddress string 	`yaml:"local_address"`

	// server 是server配置项     客户端四者都需要
	// 以上是基础信息

	AuthMethod byte			`yaml:"authority_method"`
	Users []*User			`yaml:"users"`

	Timeout int 			`yaml:"timeout"`

	// Global Plugin Config 下面两个配置，是用于握手阶段的
	ComConfig *CommConfig	`yaml:"com_config"`



	Debug bool				`yaml:"debug"`
	Slog bool				`yaml:"slog"`
}

func (config *Config) String() string {
	return fmt.Sprintf("Server=%v\nServerPort=%v\nLocalPort=%v\nLocalAddress=%v\nUsers=%v",
		config.GetServers(), config.ServerPort, config.LocalPort, config.LocalAddress, config.Users)
}

func ParseConfig(path string) (config *Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	config = &Config{}
	if err = yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}

	if config.Timeout == 0 {
		config.Timeout = 180		// 默认三分钟
	}
	if config.LocalAddress == "" {
		config.LocalAddress = "127.0.0.1"
	}
	if config.LocalPort == 0 {
		config.LocalPort = 1080
	}

	return
}

func (c *Config) GetLocalString() string {
	return c.LocalAddress
}

func (c *Config) GetLocalAddr() string {
	return fmt.Sprintf("%s:%d", c.LocalAddress, c.LocalPort)
}

func (c *Config) ServerPortString() string {
	return strconv.Itoa(c.ServerPort)
}

func (c *Config) LocalPortString() string {
	return strconv.Itoa(c.LocalPort)
}

func (config *Config) GetServers() []string {
	if config.Server == nil {
		return nil
	}

	if s, ok := config.Server.(string); ok {
		return []string{s}
	}

	if arr, ok := config.Server.([]interface{}); ok {
		serverArr := make([]string, len(arr))
		for i, s := range arr {
			if serverArr[i], ok = s.(string); !ok {
				goto typeError
			}
		}
		return serverArr
	}
typeError:
	panic(fmt.Sprintf("Config.Server type error %v", reflect.TypeOf(config.Server)))
}
