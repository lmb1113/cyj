package controller

import (
	"cyj/dto"
	"cyj/server/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Update(c *gin.Context) {
	var resp dto.CheckUpdateResp
	version, _ := strconv.Atoi(c.GetString("x-version"))
	if version < 6 {
		resp.NeedUpdate = true
		resp.Url = "https://c.0a0a.cn/cyj/1.0.6.exe"
		resp.Msg = "优化更多功能"
	} else {
		resp.NeedUpdate = false
	}
	common.RespOk(c, resp)
	return
}
