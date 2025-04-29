package proxy

import (
	"cyj/global"
	client2 "cyj/pkg/frp/client"
	"cyj/pkg/frp/cmd/frpc/sub"
	"cyj/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Options struct {
	IpAddr  string           `json:"ip_addr"`
	Service *client2.Service `json:"client"`
}

func NewProxy() *Options {
	o := &Options{
		IpAddr: "",
	}
	return o
}

func (o *Options) Run(addr string, proxyType string) (string, error) {
	if !utils.ValidateIPPortFormat(addr) {
		return "", errors.New("本地地址错误 [ip:端口]")
	}
	config := global.NewDefaultConfig()

	info := strings.Split(addr, ":")
	if len(info) != 2 {
		return "", errors.New("本地地址错误 [ip:端口]")
	}

	config.ProxyType = proxyType
	config.LocalAddr = info[0]
	config.LocalPort, _ = strconv.Atoi(info[1])
	global.SetClientInfo(config)

	service, err := sub.RunClient("")
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	o.Service = service
	return config.GetProxyUrl(), nil
}

func (o *Options) Stop() error {
	utils.Go(func() {
		o.Service.Close()
	})
	return nil
}
