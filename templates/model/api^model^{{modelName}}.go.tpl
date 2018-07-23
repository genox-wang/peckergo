package model

import (
	"fmt"
	"strconv"
	"time"

	"{{projectName}}/api/utils/json"
	"{{projectName}}/api/utils/log"

	cache "ti-ding.com/wangji/gocachemid"
)


var (
	{{modelName}}CountCache *cache.Cache
)

// {{ModelName}} 用户模型
type {{ModelName}} struct {
	Model
	// TODO
}

// Table{{ModelName}} 返回表单用户数据模型
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
	{{modelName}}CountCache.DelWithPrefix("{{projectName}}_{{modelName}}_")
	return DB.Create(m).Error
}

// Save{{ModelName}} 更新 {{ModelName}}
func Save{{ModelName}}(m *{{ModelName}}) error {
	return DB.Model(m).Updates(m).Error
}

// Delete{{ModelName}} 删除 {{ModelName}}
func Delete{{ModelName}}(id uint) error {
	m := &{{ModelName}}{}
	m.ID = id
	{{modelName}}CountCache.DelWithPrefix("{{projectName}}_{{modelName}}_")
	return DB.Delete(m).Error
}


// All{{ModelName}}s 获取所有 {{ModelName}}s
func All{{ModelName}}s(meta *TableMeta) *Table{{ModelName}} {
	metaJSON, _ := json.Marshal(meta)
	countCache, _ := {{modelName}}CountCache.Get(metaJSON)
	count, _ := strconv.ParseUint(countCache, 10, 64)

	newDB := WrapMeta(*meta, DB)
	{{modelName}}s := make([]*{{ModelName}}, 0)
	newDB.Find(&{{modelName}}s)
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
