package middleware

import (
	"peckergo/api/model"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	//JwtScretKey jwt secret key.
	JwtScretKey = "Q7nQGjXJyaRVaQvHZUPJReXs"
)

// JWTAuthRequired is auth middleware
func JWTAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.Request.Header.Get("Authorization")
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

			if d, ok := claims["id"]; ok {
				u.ID = uint(d.(float64))
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

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
		"id":      user.ID,
		"name":    user.DisplayName,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 令牌有效期为一周
		"created": time.Now().Unix(),
	})

	// log.Warn(user)

	token, err := t.SignedString([]byte(JwtScretKey))
	if err != nil {
		log.Error(err.Error())
	}
	return token
}
