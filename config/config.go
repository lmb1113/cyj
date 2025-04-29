package config

import (
	"cyj/utils"
	"embed"
	"github.com/BurntSushi/toml"
	"sync"
)

//go:embed config.toml
var ConfigFs embed.FS

const configPath = "config.toml"

var (
	cfg  *AppConfig
	once sync.Once
)

func Config() *AppConfig {
	once.Do(func() {
		data, err := ConfigFs.ReadFile(configPath)
		if err != nil {
			panic(err)
		}
		var localConfig LocalConfig
		if _, err := toml.Decode(string(data), &localConfig); err != nil {
			panic(err)
		}
		cfg = &AppConfig{
			LocalConfig: localConfig,
		}
	})
	return cfg
}

type AppConfig struct {
	LocalConfig
	RemoteConfig
}

type LocalConfig struct {
	Version    int    `json:"version"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	ApiBaseUrl string `json:"apiBaseUrl"`
}
type RemoteConfig struct {
	FrpServerConfig
}

type FrpServerConfig struct {
	ClientName   string   `json:"client_name"`
	Token        string   `json:"token"`
	FrpPort      int      `json:"frp_port"`
	ServerPort   int      `json:"server_port"`
	ServerAddr   string   `json:"server_addr"`
	HttpDomain   string   `json:"http_domain"`
	HttpsDomains []string `json:"https_domains"`
}

func (f *FrpServerConfig) GetRandomDomains() string {
	l := len(f.HttpsDomains)
	if l == 1 {
		return f.HttpsDomains[0]
	}
	if l > 1 {
		val, _ := utils.GetRandomElement(f.HttpsDomains)
		return val
	}
	return ""
}

func SetRemoteConfig(rConfig RemoteConfig) {
	Config().RemoteConfig = rConfig
	return
}
