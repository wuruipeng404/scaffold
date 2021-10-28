/*
* @Author: Rumple
* @Email: ruipeng.wu@cyclone-robotics.com
* @DateTime: 2021/9/9 10:22
 */

package scaffold

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wuruipeng404/scaffold/util"
)

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
	return []string{"Origin", "Content-Type", "Content-Length", " Accept-Encoding", "X-CSRF-VerifyToken", "Authorization"}
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

func GracefulLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {

		statusColor := params.StatusCodeColor()
		methodColor := params.MethodColor()
		resetColor := params.ResetColor()

		if params.Latency > time.Minute {
			params.Latency = params.Latency - params.Latency%time.Second
		}
		return fmt.Sprintf("[GIN] %s [%s] \"%s %s %s %s %s %d %s %s\" \"%s\" \n",
			params.TimeStamp.Format(util.TimeFormatString),
			params.ClientIP,
			methodColor, params.Method, resetColor,
			params.Path,
			statusColor, params.StatusCode, resetColor,
			params.Latency,
			params.ErrorMessage,
		)
	})
}
