package model

import (
	"fmt"
	"strconv"
	"time"

	"{{projectName}}/api/utils/json"
	
	log "github.com/sirupsen/logrus"

	cache "ti-ding.com/wangji/gocachemid"
)


var (
	{{modelName}}CountCache *cache.Cache
)

// {{ModelName}} {{ModelName}}模型
type {{ModelName}} struct {
	Model
	// TODO
}

// 分表
// TableName 设置表名
func (m {{ModelName}}) TableName() string {
	loc, err := time.LoadLocation("Asia/Chongqing")
	now := time.Now()
	if err == nil {
		now = now.In(loc)
	} else {
		log.Error("{{ModelName}} TableName: ", err.Error())
	}
	day := now.Format("060102")
	return "{{model_name}}s_" + day
}


// Table{{ModelName}} 返回表单{{ModelName}}数据模型
type Table{{ModelName}} struct {
	Data []*{{ModelName}}    `json:"data"`
	Meta *TableMeta `json:"meta"`
}

func init() {
	{{modelName}}CountCache = cache.NewCache(&cache.ClientGoCache{}, "{{modelName}}_", func(fs ...string) string {
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
		newDB.Model({{ModelName}}{}).Count(&count)
		return fmt.Sprintf("%d", count)
	}, time.Minute*5, true)
}

// New{{ModelName}} 创建 {{ModelName}}
func New{{ModelName}}(m *{{ModelName}}) error {
	err := DB.Create(m).Error
	if err == nil {
		{{modelName}}CountCache.DelWithPrefix("{{projectName}}_{{modelName}}_")
	}
	return err
}

// Save{{ModelName}} 更新 {{ModelName}}
func Save{{ModelName}}(m *{{ModelName}}) error {
	return DB.Model(m).Updates(m).Error
}

// Delete{{ModelName}} 删除 {{ModelName}}
func Delete{{ModelName}}(id uint) error {
	m := &{{ModelName}}{}
	m.ID = id
	err := DB.Delete(m).Error
	if err == nil {
		{{modelName}}CountCache.DelWithPrefix("{{projectName}}_{{modelName}}_")
	}
	return err
}


// All{{ModelName}}s 获取所有 {{ModelName}}s
func All{{ModelName}}s(meta *TableMeta) *Table{{ModelName}} {
// 分表
//func All{{ModelName}}s(meta *TableMeta, suffix string) *Table{{ModelName}} {
	countMeta := &TableMeta{
		Filter: meta.Filter,
	}
	metaJSON, _ := json.Marshal(countMeta)
	countCache, _ := {{modelName}}CountCache.Get(metaJSON)
	count, _ := strconv.ParseUint(countCache, 10, 64)

	newDB := WrapMeta(*meta, DB)
	{{modelName}}s := make([]*{{ModelName}}, 0)
	// 不分表
	newDB.Find(&{{modelName}}s)
	// 分表
	//newDB.Table(fmt.Sprintf("sdk_requests_%s", suffix)).Find(&logAuthRequests)
	meta.Pagination.Total = uint(count)
	return &Table{{ModelName}}{
		Data: {{modelName}}s,
		Meta: meta,
	}
}

// {{ModelName}}ByID 通过 id 获取 {{ModelName}}
func {{ModelName}}ByID(id uint) *{{ModelName}} {
	var m {{ModelName}}
	DB.Where("id = ?", id).First(&m)
	return &m
}


// All{{ModelName}}IDNameMap 获取所有 Server ID-Name 映射
func All{{ModelName}}IDNameMap() map[uint]string {
	var ms []*{{ModelName}}
	if DB.Select("id, name").Find(&ms).Error != nil {
		return map[uint]string{}
	}
	mMap := make(map[uint]string, 0)
	for _, m := range ms {
		mMap[m.ID] = m.Name
	}
	return mMap
}
