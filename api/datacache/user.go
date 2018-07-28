package datacache

import (
	"fmt"
	"strconv"
)

const (
	userPrefix = "peckergo_pwc"
)

// SetPwChangeTime 设置用户密码修改修改
func SetPwChangeTime(userID uint, changeTime int64) {
	Set(fmt.Sprintf("%s_%d", userPrefix, userID), strconv.FormatInt(changeTime, 10), 0)
}

// DelPwChangeTime 删除用户密码修改实践
func DelPwChangeTime(userID uint) {
	Del(fmt.Sprintf("%s_%d", userPrefix, userID))
}

// GetPwChangeTime 获取用户密码修改时间
func GetPwChangeTime(userID uint) (int64, error) {
	timeVal, err := Get(fmt.Sprintf("%s_%d", userPrefix, userID))
	if err != nil {
		return 0, err
	}
	changeTime, err := strconv.ParseInt(timeVal, 10, 64)
	if err != nil {
		return 0, err
	}
	return changeTime, nil
}
