package controller

import (
	"console-template/api/middleware"
	"console-template/api/model"
	"console-template/api/utils/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// LoginPost is a function
func LoginPost(c *gin.Context) {
	var user model.User

	if err := json.BindGinJSON(c, &user); err == nil {
		if err := model.Login(&user); err == nil {
			json.WriteGinJSON(c, http.StatusOK, gin.H{
				"token": middleware.GetJWTToken(user),
			})
		} else {
			json.WriteGinJSON(c, http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
		}
		return
	}

	json.WriteGinJSON(c, http.StatusBadRequest, gin.H{
		"msg": "param error!",
	})
}

// NewUserPost create user
func NewUserPost(c *gin.Context) {
	var user model.User

	if err := json.BindGinJSON(c, &user); err == nil {
		if err := model.NewUser(&user); err == nil {
			json.WriteGinJSON(c, http.StatusOK, gin.H{})
		} else {
			json.WriteGinJSON(c, http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
		return
	}

	json.WriteGinJSON(c, http.StatusBadRequest, gin.H{
		"msg": "param error!",
	})
}

// AllUsersGet get all user
func AllUsersGet(c *gin.Context) {
	meta := model.TableMetaFromQuery(c)
	json.WriteGinJSON(c, http.StatusOK, model.AllUsers(meta))
}

// UpdateUserPut 更新 User
func UpdateUserPut(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	m := &model.User{}

	if err := json.BindGinJSON(c, m); err != nil {
		json.WriteGinJSON(c, http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	m.ID = uint(id)

	if err := model.SaveUser(m); err != nil {
		json.WriteGinJSON(c, http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	json.WriteGinJSON(c, http.StatusOK, gin.H{})
}

// UserDelete 更新 User
func UserDelete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if err := model.DeleteUser(uint(id)); err != nil {
		json.WriteGinJSON(c, http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	json.WriteGinJSON(c, http.StatusOK, gin.H{})
}
