/*
* @Author: Rumple
* @Email: ruipeng.wu@cyclone-robotics.com
* @DateTime: 2021/8/23 14:48
 */

package control

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors 处理跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-VerifyToken, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
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
