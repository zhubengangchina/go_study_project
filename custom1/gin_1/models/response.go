package models

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`    // 0 表示成功，非 0 表示失败
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据
}

const (
	SUCCESS = 0
	ERROR   = 1
)

func OK(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code:    SUCCESS,
		Message: "成功",
		Data:    data,
	})
}

func Fail(c *gin.Context, msg string) {
	c.JSON(200, Response{
		Code:    ERROR,
		Message: msg,
		Data:    nil,
	})
}
