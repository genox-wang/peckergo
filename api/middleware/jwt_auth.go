package middleware

import (
	"console-template/api/model"
	"net/http"

	"console-template/api/utils/log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	//JwtScretKey jwt secret key.
	JwtScretKey = "JWT_SCRET_KEY"
)

// JWTAuthRoleRequired 用户角色认证中间件
func JWTAuthRoleRequired(roles ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, _ := c.Get("user")
		user := u.(model.User)
		log.Warn(user)
		for _, r := range roles {
			if r == user.Role {
				c.Next()
				return
			}
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

// JWTAuthRequired is auth middleware
func JWTAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.Request.Header.Get("Authorization")
		log.Info(tokenString)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtScretKey), nil
		})

		if err != nil {
			log.Error("Parse => ", err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			u := model.User{}

			if d, ok := claims["name"]; ok {
				u.DisplayName = d.(string)
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			if r, ok := claims["role"]; ok {
				u.Role = int(r.(float64))
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			if d, ok := claims["id"]; ok {
				u.ID = uint(d.(float64))
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			c.Set("user", u)
			c.Next()
			return
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

// GetJWTToken build a jwt token for user
func GetJWTToken(user model.User) (token string) {

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"name": user.DisplayName,
		"role": user.Role,
		// "exp":  time.Now().Add(time.Minute * 5).Unix(),
		"exp": 0,
	})

	log.Warn(user)

	token, err := t.SignedString([]byte(JwtScretKey))
	if err != nil {
		log.Error(err.Error())
	}
	return token
}
