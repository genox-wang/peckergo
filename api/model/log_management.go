package model

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"peckergo/api/utils/json"

	log "github.com/sirupsen/logrus"

	cache "ti-ding.com/wangji/gocachemid"
)

var (
	logManagementCountCache *cache.Cache
)

// LogManagement LogManagement模型
type LogManagement struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Path      string    `json:"path"`
	Method    string    `json:"method"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

// 分表取消下面方法注释
// TableName 设置表名
// func (m LogManagement) TableName() string {
// 	loc, err := time.LoadLocation("Asia/Chongqing")
// 	now := time.Now()
// 	if err == nil {
//		now = now.In(loc)
// 	} else {
// 		log.Error("LogManagement TableName: ", err.Error())
// 	}
//	day := now.Format("060102")
//	return "log_managements_" + day
// }

// TableLogManagement 返回表单LogManagement数据模型
type TableLogManagement struct {
	Data []*LogManagement `json:"data"`
	Meta *TableMeta       `json:"meta"`
}

func init() {
	logManagementCountCache = cache.NewCache(&cache.ClientGoCache{}, "peckergo_log_management_cnt_", func(fs ...string) (string, error) {
		if len(fs) < 1 {
			return "0", errors.New("len(fs) < 1")
		}
		var meta *TableMeta
		err := json.UnmarshalFromString(fs[0], &meta)
		if err != nil {
			log.Error(err.Error())
			return "0", errors.New(err.Error())
		}
		newDB := WrapMeta(*meta, DB)
		var count uint
		newDB.Model(LogManagement{}).Count(&count)
		return fmt.Sprintf("%d", count), nil
	}, time.Minute*5, true)
}

// NewLogManagement 创建 LogManagement
func NewLogManagement(m *LogManagement) error {
	err := DB.Create(m).Error
	if err == nil {
		logManagementCountCache.DelWithPrefix("peckergo_log_management_cnt_")
	}
	return err
}

// AllLogManagements 获取所有 LogManagements
func AllLogManagements(meta *TableMeta) *TableLogManagement {
	// 分表注释上行代码，取消注释下行代码
	//func AllLogManagements(meta *TableMeta, suffix string) *TableLogManagement {
	countMeta := &TableMeta{
		Filter: meta.Filter,
	}
	metaJSON, _ := json.MarshalToString(countMeta)
	countCache, _ := logManagementCountCache.Get(metaJSON)
	count, _ := strconv.ParseUint(countCache, 10, 64)

	newDB := WrapMeta(*meta, DB)
	logManagements := make([]*LogManagement, 0)
	// 分表注释下行代码
	newDB.Find(&logManagements)
	// 分表取消注释下行代码
	//newDB.Table(fmt.Sprintf("log_managements_%s", suffix)).Find(&logManagements)
	meta.Pagination.Total = uint(count)
	return &TableLogManagement{
		Data: logManagements,
		Meta: meta,
	}
}
