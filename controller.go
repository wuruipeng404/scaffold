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
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    any    `json:"data,omitempty"`
	ApiPage `json:",inline"`
}

type BeautyController struct{}

// func (c *BeautyController) TraceId(ctx *gin.Context) string {
// 	return ctx.GetHeader(_TraceKey)
// }

// OK response with data
func (c *BeautyController) OK(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, ApiResponse{Code: _Success, Msg: _SuccessMsg, Data: data})
}

// PageOk response with page data
func (c *BeautyController) PageOk(ctx *gin.Context, data any, page ApiPage) {
	ctx.JSON(http.StatusOK, ApiResponse{
		Code:    _Success,
		Msg:     _SuccessMsg,
		Data:    data,
		ApiPage: page,
	})
}

// RawOK without code msg
func (c *BeautyController) RawOK(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}

// PureOK only http code 200
func (c *BeautyController) PureOK(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

// Failed with IError
func (c *BeautyController) Failed(ctx *gin.Context, ie er.IError) {
	ctx.JSON(http.StatusOK, ApiResponse{Code: ie.Code(), Msg: ie.Message()})
}

// FailedC with custom code and message
func (c *BeautyController) FailedC(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, ApiResponse{Code: code, Msg: message})
}

// FailedD with IError and Data
func (c *BeautyController) FailedD(ctx *gin.Context, ie er.IError, data any) {
	ctx.JSON(http.StatusOK, ApiResponse{Code: ie.Code(), Msg: ie.Message(), Data: data})
}

// FailedE with default code and raw error
func (c *BeautyController) FailedE(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, ApiResponse{Code: _Failure, Msg: err.Error()})
}

// FailedDyn join IError message adn dynamic raw error message
func (c *BeautyController) FailedDyn(ctx *gin.Context, ie er.IError, err error) {
	ctx.JSON(http.StatusOK, ApiResponse{Code: ie.Code(), Msg: fmt.Sprintf("%s %s", ie.Message(), err)})
}
