package net

import "github.com/gin-gonic/gin"

// 获取真实 IP
func GetRealIP(c *gin.Context) string {
	var ip string
	if c.GetHeader("X-Forwarded-For") == "" {
		ip = c.Request.RemoteAddr
	} else {
		ip = c.GetHeader("X-Forwarded-For")
	}
	return ip
}
