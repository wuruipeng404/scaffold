/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/9/9 10:21
 */

package scaffold

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

func (c *BeautyController) OK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, ApiResponse{Code: _Success, Msg: _SuccessMsg, Data: data})
}

// RawOK without code msg
func (c *BeautyController) RawOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func (c *BeautyController) PureOK(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func (c *BeautyController) Failed(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, ApiResponse{Code: _Failure, Msg: err.Error()})
}

func (c *BeautyController) FailedWithError(ctx *gin.Context, err IError) {
	ctx.JSON(http.StatusOK, ApiResponse{Code: err.Code(), Msg: err.Message()})
}
