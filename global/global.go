package global

import (
	"cyj/utils"
	"fmt"
	"sync"
)

const remotePort = 7777
const Version = 1

type Config struct {
	LocalAddr   string
	LocalPort   int
	RemotePort  int
	RemoteAddr  string
	ServicePort int
	ServiceAddr string
	Token       string
	Name        string
}

func (c *Config) GetProxyUrl() string {
	return fmt.Sprintf("%s:%d", c.RemoteAddr, c.RemotePort)
}

var once sync.Once

func NewDefaultConfig() *Config {
	once.Do(func() {
		gConfig = &Config{
			ServicePort: remotePort,
			ServiceAddr: "c.0a0a.cn",
			Token:       "12345",
			RemoteAddr:  "c.0a0a.cn",
			Name:        utils.GenerateName(),
		}
	})
	return gConfig
}

var gConfig *Config

func SetConfig(config *Config) {
	gConfig = config
}

func GetConfig() *Config {
	return gConfig
}
