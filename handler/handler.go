package handler

import (
	"github.com/XGHXT/SYOJ-Backend/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendBadRequest(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusBadRequest, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendError(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusInternalServerError, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendForbidden(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusForbidden, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}