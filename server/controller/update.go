package controller

import (
	"cyj/dto"
	"cyj/server/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

const lastVersion = "1.0.7"

func Update(c *gin.Context) {
	var resp dto.CheckUpdateResp
	version, _ := strconv.Atoi(c.GetString("x-version"))
	if version < 7 {
		resp.NeedUpdate = true
		resp.Url = "https://c.0a0a.cn/cyj/" + lastVersion + ".exe"
		resp.Msg = "优化更多功能"
	} else {
		resp.NeedUpdate = false
	}
	common.RespOk(c, resp)
	return
}

func Download(c *gin.Context) {
	c.Redirect(302, "https://c.0a0a.cn/cyj/"+lastVersion+".exe")
	return
}
