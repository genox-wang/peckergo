package model

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
)

const (
	// RoleAdmin 管理员角色
	RoleAdmin = 1
	// RoleOperator 操作员角色
	RoleOperator = 2
)

// User 用户模型
type User struct {
	Model
	DisplayName string `json:"display_name"`
	Username    string `json:"username" gorm:"unique;not null"`
	Password    string `json:"password" gorm:"not null"`
	Role        int    `json:"role" gorm:"default:0"`
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

// AllUsers 获取所有 Users
func AllUsers(meta *TableMeta) *TableUser {
	newDB := WrapMeta(*meta, DB)
	var count uint
	users := make([]*User, 0)
	newDB.Find(&users).Count(&count)
	meta.Pagination.Total = count
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
