package controller

import (
	"net/http"
	"peckergo/api/middleware"
	"peckergo/api/model"
	ginutils "peckergo/api/utils/gin"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// CaptchaGet 获取图形验证码
func CaptchaGet(c *gin.Context) {
	//config struct for digits
	//数字验证码配置
	var configD = base64Captcha.ConfigDigit{
		Height:     50,
		Width:      150,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 4,
	}

	//create a digits captcha.
	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//write to base64 string.
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)

	ginutils.WriteGinJSON(c, http.StatusOK, gin.H{
		"key":     idKeyD,
		"captcha": base64stringD,
	})
}

// LoginPost is a function
func LoginPost(c *gin.Context) {
	var user model.User

	if err := ginutils.BindGinJSON(c, &user); err == nil {
		verifyResult := base64Captcha.VerifyCaptcha(user.CaptchaKey, user.Captcha)
		if !verifyResult {
			ginutils.WriteGinJSON(c, http.StatusUnauthorized, gin.H{
				"msg": "验证码错误",
			})
			return
		}
		if err := model.Login(&user); err == nil {
			ginutils.WriteGinJSON(c, http.StatusOK, gin.H{
				"token": middleware.GetJWTToken(user),
			})
		} else {
			ginutils.WriteGinJSON(c, http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
		}
		return
	}

	ginutils.WriteGinJSON(c, http.StatusBadRequest, gin.H{
		"msg": "传参数错误!",
	})
}

// NewUserPost create user
func NewUserPost(c *gin.Context) {
	var user model.User

	if err := ginutils.BindGinJSON(c, &user); err == nil {
		if err := model.NewUser(&user); err == nil {
			ginutils.WriteGinJSON(c, http.StatusOK, gin.H{})
		} else {
			ginutils.WriteGinJSON(c, http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
		return
	}

	ginutils.WriteGinJSON(c, http.StatusBadRequest, gin.H{
		"msg": "传参数错误!",
	})
}

// AllUsersGet get all user
func AllUsersGet(c *gin.Context) {
	meta := model.TableMetaFromQuery(c)
	ginutils.WriteGinJSON(c, http.StatusOK, model.AllUsers(meta))
}

// UserByIDGet get user by id
func UserByIDGet(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	log.Info("UserByIDGet ", id)
	m := model.UserByID(uint(id))
	ginutils.WriteGinJSON(c, http.StatusOK, m)
}

// UpdateUserPut 更新 User
func UpdateUserPut(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	m := &model.User{}

	if err := ginutils.BindGinJSON(c, m); err != nil {
		ginutils.WriteGinJSON(c, http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	m.ID = uint(id)

	if err := model.SaveUser(m); err != nil {
		ginutils.WriteGinJSON(c, http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ginutils.WriteGinJSON(c, http.StatusOK, gin.H{})
}

// UserDelete 更新 User
func UserDelete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if err := model.DeleteUser(uint(id)); err != nil {
		ginutils.WriteGinJSON(c, http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ginutils.WriteGinJSON(c, http.StatusOK, gin.H{})
}

// AllUserIDNameMapGet 获得所有 User ID-Name 映射
func AllUserIDNameMapGet(c *gin.Context) {
	mMap := model.AllUserIDNameMap()
	ginutils.WriteGinJSON(c, http.StatusOK, mMap)
}
