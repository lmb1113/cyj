package global

import (
	"cyj/utils"
	"fmt"
	"sync"
	"time"
)

const remotePort = 7777
const Version = 6

type Config struct {
	LocalAddr     string
	LocalPort     int
	RemotePort    int
	RemoteAddr    string
	ServicePort   int
	ServiceAddr   string
	Token         string
	Name          string
	ProxyType     string
	SubDomain     string
	CustomDomains []string
}

func (c *Config) GetProxyUrl() string {
	if c.ProxyType == "https" {
		return fmt.Sprintf("https://%s", c.CustomDomains[0])
	}
	return fmt.Sprintf("http://%s:%d", c.RemoteAddr, c.RemotePort)
}

var once sync.Once

func NewDefaultConfig() *Config {
	subDomain := utils.RandomString(8, time.Now().UnixMicro())
	once.Do(func() {
		gConfig = &Config{
			ServicePort:   remotePort,
			ServiceAddr:   "c.0a0a.cn",
			Token:         "12345",
			RemoteAddr:    "c.0a0a.cn",
			Name:          utils.GenerateName(),
			CustomDomains: []string{subDomain + ".c.0a0a.cn"},
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
