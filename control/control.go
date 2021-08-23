/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/8/23 14:28
 */

package control

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	_Success = iota
	_Failure

	_SuccessMsg = "success"
)

type ApiResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type BeautyController struct{}

func NewBeautyControl() *BeautyController {
	return new(BeautyController)
}

func (*BeautyController) OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ApiResponse{Code: _Success, Msg: _SuccessMsg, Data: data})
}

// RawOK without code msg
func (*BeautyController) RawOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func (*BeautyController) PureOK(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (*BeautyController) Error(c *gin.Context, err error) {
	c.JSON(http.StatusOK, ApiResponse{Code: _Failure, Msg: err.Error()})
}

func (*BeautyController) ErrorWithCode(c *gin.Context, code int, err error) {
	c.JSON(http.StatusOK, ApiResponse{Code: code, Msg: err.Error()})
}
