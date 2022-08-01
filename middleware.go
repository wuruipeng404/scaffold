/*
* @Author: Rumple
* @Email: wrp357711589@gmail.com
* @DateTime: 2021/9/9 10:22
 */

package scaffold

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wuruipeng404/scaffold/util"
)

const _TraceKey = "trace-id"

func DefaultCorsMap() map[string]string {
	return map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Max-Age":           "86400",
		"Access-Control-Allow-Methods":     "POST, GET, OPTIONS, PUT, DELETE, PATCH",
		"Access-Control-Allow-Headers":     strings.Join(DefaultAllowHeaders(), ","),
		"Access-Control-Expose-Headers":    "Content-Length",
		"Access-Control-Allow-Credentials": "true",
	}
}

func DefaultAllowHeaders() []string {
	return []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-VerifyToken", "Authorization"}
}

// Cors 处理跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		for k, v := range DefaultCorsMap() {
			c.Writer.Header().Set(k, v)
		}
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
	}
}

func NewCors(headers map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for k, v := range headers {
			c.Writer.Header().Set(k, v)
		}
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
	}
}

func NewCorsWithDefault(headers map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for k, v := range DefaultCorsMap() {
			c.Writer.Header().Set(k, v)
		}

		for k, v := range headers {
			c.Writer.Header().Set(k, v)
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
	}
}

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetHeader(_TraceKey)
		if traceId == "" {
			traceId = genTraceId()
		}
		c.Set(_TraceKey, traceId)
		c.Header("R-Request-Trace", traceId)
		c.Next()
	}
}

func genTraceId() string {
	return uuid.NewString()
}

func GracefulLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {

		statusColor := params.StatusCodeColor()
		methodColor := params.MethodColor()
		resetColor := params.ResetColor()

		if params.Latency > time.Minute {
			params.Latency = params.Latency - params.Latency%time.Second
		}
		return fmt.Sprintf("%v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
			params.TimeStamp.Format(util.TimeFormatMill),
			statusColor, params.StatusCode, resetColor,
			params.Latency,
			params.ClientIP,
			methodColor, params.Method, resetColor,
			params.Path,
			params.ErrorMessage,
		)
	})
}
