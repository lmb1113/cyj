package global

import (
	"cyj/config"
	"fmt"
	"sync"
)

type ClineInfo struct {
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

func (c *ClineInfo) GetProxyUrl() string {
	if c.ProxyType == "https" {
		return fmt.Sprintf("https://%s", c.CustomDomains[0])
	}
	return fmt.Sprintf("http://%s:%d", c.RemoteAddr, c.RemotePort)
}

var once sync.Once

func NewDefaultConfig() *ClineInfo {
	once.Do(func() {
		gConfig = &ClineInfo{
			ServicePort:   config.Config().ServerPort,
			ServiceAddr:   config.Config().ServerAddr,
			Token:         config.Config().Token,
			RemoteAddr:    config.Config().HttpDomain,
			Name:          config.Config().ClientName,
			CustomDomains: config.Config().HttpsDomains,
			RemotePort:    config.Config().FrpPort,
		}
	})
	return gConfig
}

var gConfig *ClineInfo

func SetClientInfo(config *ClineInfo) {
	gConfig = config
}

func GetClientInfo() *ClineInfo {
	return gConfig
}
