package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type baseHandler struct {
	Code int         `json:"code" binding:"required"`
	Msg  string      `json:"msg" binding:"required"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &baseHandler{
		Code: 0,
		Msg:  "ok",
		Data: data,
	})
}

func ResponseError(c *gin.Context, errMsg string, errCode int) {
	c.JSON(http.StatusOK, &baseHandler{
		Code: errCode,
		Msg:  errMsg,
	})
}
