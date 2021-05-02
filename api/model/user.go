package model

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"peckergo/api/utils/json"

	log "github.com/sirupsen/logrus"

	cache "github.com/wilfordw/go-cachemid"
)

const (
	// RoleAdmin 管理员角色
	RoleAdmin = 1
	// RoleClient 客户角色
	RoleClient = 2
	// Salt 为密码加密加盐
	Salt = "Pn3dQvLM"
)

var (
	userCountCache *cache.Cache
)

// User User模型
type User struct {
	Model
	DisplayName string `json:"display_name"`
	Username    string `json:"username" gorm:"unique;not null"`
	Password    string `json:"password" gorm:"not null"`
	Role        int    `json:"role" gorm:"default:0"` // 1 管理员  2 客户
	CaptchaKey  string `json:"captcha_key" gorm:"-"`
	Captcha     string `json:"captcha" gorm:"-"`
}

// TableUser 返回表单User数据模型
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
	password += Salt
	b := sha1.New()
	io.WriteString(b, password)
	return fmt.Sprintf("%x", b.Sum(nil))
}

func init() {
	funcReadData := func(fs ...string) (string, bool, error) {
		if len(fs) < 1 {
			return "0", true, errors.New("len(fs) < 1")
		}
		var meta *TableMeta
		err := json.UnmarshalFromString(fs[0], &meta)
		if err != nil {
			log.Error(err.Error())
			return "0", true, errors.New(err.Error())
		}
		newDB := WrapMeta(*meta, DB)
		var count uint
		newDB.Model(User{}).Count(&count)
		return fmt.Sprintf("%d", count), true, nil
	}

	userCountCache = &cache.Cache{
		CacheClient:      cache.NewGoCache(),
		KeyPrefix:        "user_",
		FuncReadData:     funcReadData,
		ExpireTime:       time.Minute * 5,
		Cache2Enabled:    true,
		Cache2ExpireTime: cache.DefaultCache2ExpirePadding,
	}
}

// AllUsers 获取所有 Users
func AllUsers(meta *TableMeta) *TableUser {
	countMeta := &TableMeta{
		Filter: meta.Filter,
	}
	metaJSON, _ := json.MarshalToString(countMeta)
	countCache, _, _ := userCountCache.Get(metaJSON)
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

// AllUserIDNameMap 获取所有 User ID-Name 映射
func AllUserIDNameMap() map[uint]string {
	var ms []*User
	if DB.Select("id, display_name").Find(&ms).Error != nil {
		return map[uint]string{}
	}
	mMap := make(map[uint]string, 0)
	for _, m := range ms {
		mMap[m.ID] = m.DisplayName
	}
	return mMap
}
