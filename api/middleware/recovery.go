package middleware

import (
	logutils "peckergo/api/utils/log"

	"github.com/gin-gonic/gin"
)

// Recovery 抓 Panic 日志
func Recovery(roles ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer logutils.PanicRecover()
		c.Next()
	}
}
