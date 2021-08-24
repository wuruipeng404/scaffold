/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/8/23 14:28
 */

package control

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wuruipeng404/scaffold/logger"
	"go.uber.org/zap"
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

type BeautyController struct {
	*zap.SugaredLogger
}

func NewBeautyControl(log *zap.SugaredLogger) *BeautyController {
	return &BeautyController{log}
}

func (c *BeautyController) OK(ctx *gin.Context, data interface{}) {
	logger.Info(data)
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
	c.Errorf("Request [ Failed ] >> %s", err)
	logger.Error(err)
	ctx.JSON(http.StatusOK, ApiResponse{Code: _Failure, Msg: err.Error()})
}

func (c *BeautyController) FailedWithCode(ctx *gin.Context, code int, err error) {
	logger.Debug(err)
	c.Errorf("Request [ FailedWithCode ] >> %d : %s", code, err)
	ctx.JSON(http.StatusOK, ApiResponse{Code: code, Msg: err.Error()})
}
