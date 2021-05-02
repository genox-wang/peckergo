package middleware

import (
	"peckergo/api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthRoleRequired 用户角色认证中间件
func AuthRoleRequired(roles ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, _ := c.Get("user")
		user := u.(model.User)

		if user.Role != model.RoleAdmin {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// log.Warn(user)
		for _, r := range roles {
			if r == user.Role {
				c.Next()
				return
			}
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
