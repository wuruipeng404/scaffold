/*
* @Author: Rumple
* @Email: ruipeng.wu@cyclone-robotics.com
* @DateTime: 2021/8/23 14:48
 */

package control

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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
