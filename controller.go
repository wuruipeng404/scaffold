/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/9/9 10:21
 */

package scaffold

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wuruipeng404/scaffold/er"
)

const (
	_Success = iota
	_Failure

	_SuccessMsg = "success"
)

type ApiPage struct {
	TotalCount int64  `json:"total_count,omitempty"` // 列表数据总数
	TotalPage  int64  `json:"total_page,omitempty"`  // 列表页总数
	PageCount  int64  `json:"page_count,omitempty"`  // 当前页数量
	PageIndex  int64  `json:"page_index,omitempty"`  // 当前页码
	Sort       string `json:"sort,omitempty"`        // 排序
}

type ApiResponse struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
	ApiPage `json:",inline"`
}

type BeautyController struct{}

func (c *BeautyController) OK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, ApiResponse{Code: _Success, Msg: _SuccessMsg, Data: data})
}

func (c *BeautyController) PageOk(ctx *gin.Context, data interface{}, page ApiPage) {
	ctx.JSON(http.StatusOK, ApiResponse{
		Code:    _Success,
		Msg:     _SuccessMsg,
		Data:    data,
		ApiPage: page,
	})
}

// RawOK without code msg
func (c *BeautyController) RawOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func (c *BeautyController) PureOK(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func (c *BeautyController) Failed(ctx *gin.Context, ie er.IError) {
	ctx.JSON(http.StatusOK, ApiResponse{Code: ie.Code(), Msg: ie.Message()})
}

func (c *BeautyController) FailedC(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, ApiResponse{Code: code, Msg: message})
}

func (c *BeautyController) FailedD(ctx *gin.Context, ie er.IError, data interface{}) {
	ctx.JSON(http.StatusOK, ApiResponse{Code: ie.Code(), Msg: ie.Message(), Data: data})
}

func (c *BeautyController) FailedE(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, ApiResponse{Code: _Failure, Msg: err.Error()})
}

func (c *BeautyController) FailedDyn(ctx *gin.Context, ie er.IError, err error) {
	ctx.JSON(http.StatusOK, ApiResponse{Code: ie.Code(), Msg: fmt.Sprintf("%s %s", ie.Message(), err)})
}
