package common

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Data any    `json:"data"`
}

// 给设备端解析用的
type Response2 struct {
	Code int             `json:"code"`
	Msg  string          `json:"message"`
	Data json.RawMessage `json:"data"`
}

const (
	CodeOk  = 0
	CodeErr = 500
)

func RespOk(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code: CodeOk,
		Data: data,
	})
	return
}

func RespFail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: CodeErr,
		Msg:  msg,
	})
	return
}
