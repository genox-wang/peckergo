package model

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"console-template/api/utils/json"
	"console-template/api/utils/log"

	cache "ti-ding.com/wangji/gocachemid"
)

const (
	// RoleAdmin 管理员角色
	RoleAdmin = 1
	// RoleOperator 操作员角色
	RoleOperator = 2
)

var (
	userCountCache *cache.Cache
)

// User 用户模型
type User struct {
	Model
	DisplayName string `json:"display_name"`
	Username    string `json:"username" gorm:"unique;not null"`
	Password    string `json:"password" gorm:"not null"`
	Role        int    `json:"role" gorm:"default:0"`
	CaptchaKey  string `json:"captcha_key" gorm:"-"`
	Captcha     string `json:"captcha" gorm:"-"`
}

// TableUser 返回表单用户数据模型
type TableUser struct {
	Data []*User    `json:"data"`
	Meta *TableMeta `json:"meta"`
}

// Login 用户登陆验证
func Login(m *User) error {
	var tmpU User
	if DB.Where("username=?", m.Username).First(&tmpU).Error != nil {
		return errors.New("Username doesn't exist")
	}
	if tmpU.Password == CrptoPassword(m.Password) {
		m.DisplayName = tmpU.DisplayName
		m.ID = tmpU.ID
		m.Role = tmpU.Role
		return nil
	}
	return errors.New("Password is incorrent")
}

// NewUser 创建 User
func NewUser(m *User) error {
	m.Password = CrptoPassword(m.Password)
	userCountCache.DelWithPrefix("user_")
	return DB.Create(m).Error
}

// SaveUser 更新 User
func SaveUser(m *User) error {
	m.Password = CrptoPassword(m.Password)
	return DB.Model(m).Updates(m).Error
}

// DeleteUser 删除 User
func DeleteUser(id uint) error {
	m := &User{}
	m.ID = id
	userCountCache.DelWithPrefix("user_")
	return DB.Delete(m).Error
}

// CrptoPassword crpto password
func CrptoPassword(password string) string {
	if password == "" {
		return ""
	}
	b := sha1.New()
	io.WriteString(b, password)
	return fmt.Sprintf("%x", b.Sum(nil))
}

func init() {
	userCountCache = cache.NewCache(&cache.ClientGoCache{}, "user_", func(fs ...string) string {
		if len(fs) < 1 {
			return "0"
		}
		var meta *TableMeta
		err := json.Unmarshal(fs[0], &meta)
		if err != nil {
			log.Error(err.Error())
			return "0"
		}
		newDB := WrapMeta(*meta, DB)
		var count uint
		newDB.Model(User{}).Count(&count)
		return fmt.Sprintf("%d", count)
	}, time.Minute*5, true)
}

// AllUsers 获取所有 Users
func AllUsers(meta *TableMeta) *TableUser {
	countMeta := &TableMeta{
		Filter: meta.Filter,
	}
	metaJSON, _ := json.Marshal(countMeta)
	countCache, _ := userCountCache.Get(metaJSON)
	count, _ := strconv.ParseUint(countCache, 10, 64)

	newDB := WrapMeta(*meta, DB)
	users := make([]*User, 0)
	newDB.Find(&users)
	meta.Pagination.Total = uint(count)
	return &TableUser{
		Data: users,
		Meta: meta,
	}
}

// UserByID 通过 id 获取 User
func UserByID(id uint) *User {
	var m User
	DB.Where("id = ?", id).First(&m)
	return &m
}
