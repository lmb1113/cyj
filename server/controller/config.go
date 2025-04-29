package controller

import (
	config2 "cyj/config"
	"cyj/server/common"
	"cyj/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func Config(c *gin.Context) {
	common.RespOk(c, &config2.RemoteConfig{
		FrpServerConfig: config2.FrpServerConfig{
			ClientName:   utils.GenerateName(),
			Token:        "12345",
			ServerPort:   7777,
			FrpPort:      utils.GenerateRemotePort(),
			ServerAddr:   "c.0a0a.cn",
			HttpDomain:   "c.0a0a.cn",
			HttpsDomains: []string{utils.RandomString(8, time.Now().UnixMicro()) + ".c.0a0a.cn"},
		},
	})
	return
}
