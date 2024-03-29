package model

import (
	"fmt"
	"strconv"
	"errors"
	"time"

	"{{projectName}}/api/utils/json"
	
	log "github.com/sirupsen/logrus"

	cache "github.com/genox-wang/go-cachemid"
)

var (
	{{modelName}}CountCache *cache.Cache
)

// {{ModelName}} {{ModelName}}模型
type {{ModelName}} struct {
	Model
	// TODO 添加模型必要字段
}

// TODO 分表取消下面方法注释
// TableName 设置表名
// func (m {{ModelName}}) TableName() string {
// 	loc, err := time.LoadLocation("Asia/Chongqing")
// 	now := time.Now()
// 	if err == nil {
//		now = now.In(loc)
// 	} else {
// 		log.Error("{{ModelName}} TableName: ", err.Error())
// 	}
//	day := now.Format("060102")
//	return "{{model_name}}s_" + day
// }


// Table{{ModelName}} 返回表单{{ModelName}}数据模型
type Table{{ModelName}} struct {
	Data []*{{ModelName}}    `json:"data"`
	Meta *TableMeta `json:"meta"`
}

func init() {
	funcReadData := func(fs ...string) (string, bool, error) {
		// TODO 分表注释下面3行代码
		if len(fs) < 1 {
			return "0", true, errors.New("len(fs) < 1")
		}
			// TODO 分表取消下面方法注释
		// if len(fs) < 2 {
		//	return "0", errors.New("len(fs) < 2")
		// }
		var meta *TableMeta
		err := json.UnmarshalFromString(fs[0], &meta)
		if err != nil {
			log.Error(err.Error())
			return "0", true, errors.New(err.Error())
		}
		newDB := WrapMeta(*meta, DB)
		var count uint	
		// TODO 分表注释下行代码
		newDB.Model({{ModelName}}{}).Count(&count)
		// TODO 分表取消下行注释
		// newDB.Table(fmt.Sprintf("{{model_name}}_%s", fs[1])).Count(&count)
		return fmt.Sprintf("%d", count), true, nil
	}

	{{modelName}}CountCache = &cache.Cache{
		CacheClient:      cache.NewGoCache(),
		KeyPrefix:        "{{projectName}}_{{model_name}}_cnt_",
		FuncReadData:     funcReadData,
		ExpireTime:       time.Minute * 5,
		Cache2Enabled:    true,
		Cache2ExpireTime: cache.DefaultCache2ExpirePadding,
	}
	
}

// New{{ModelName}} 创建 {{ModelName}}
func New{{ModelName}}(m *{{ModelName}}) error {
	// TODO 分表取消注释下行代码
	//if !HasTable(&{{ModelName}}{}) {
	//	DB.CreateTable(&{{ModelName}}{})
	//}
	err := DB.Create(m).Error
	if err == nil {
		{{modelName}}CountCache.DelWithPrefix("{{projectName}}_{{model_name}}_cnt_")
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
		{{modelName}}CountCache.DelWithPrefix("{{projectName}}_{{model_name}}_cnt_")
	}
	return err
}


// All{{ModelName}}s 获取所有 {{ModelName}}s
func All{{ModelName}}s(meta *TableMeta) *Table{{ModelName}} {
// TODO 分表注释上行代码，取消注释下行代码
//func All{{ModelName}}s(meta *TableMeta, suffix string) *Table{{ModelName}} {
	countMeta := &TableMeta{
		Filter: meta.Filter,
	}
	metaJSON, _ := json.MarshalToString(countMeta)
	// TODO 分表注释下行
	countCache, _, _ := {{modelName}}CountCache.Get(metaJSON)
	// TODO 分表取消下行注释
	// countCache, _, _ := logAdRequestCountCache.Get(metaJSON, suffix)
	count, _ := strconv.ParseUint(countCache, 10, 64)

	newDB := WrapMeta(*meta, DB)
	{{modelName}}s := make([]*{{ModelName}}, 0)
	// TODO 分表注释下行代码
	newDB.Find(&{{modelName}}s)
	// TODO 分表取消注释下行代码
	//newDB.Table(fmt.Sprintf("{{model_name}}s_%s", suffix)).Find(&{{modelName}}s)
	meta.Pagination.Total = uint(count)
	return &Table{{ModelName}}{
		Data: {{modelName}}s,
		Meta: meta,
	}
}

// {{ModelName}}ByID 通过 id 获取 {{ModelName}}
func {{ModelName}}ByID(id uint) *{{ModelName}} {
	var m {{ModelName}}
	if DB.Where("id = ?", id).First(&m).Error != nil {
		return nil
	}
	return &m
}

// TODO 为前端暴露 ID-Name 映射
// All{{ModelName}}IDNameMap 获取所有 Server ID-Name 映射
// func All{{ModelName}}IDNameMap() map[uint]string {
// 	 var ms []*{{ModelName}}
// 	 if DB.Select("id, name").Find(&ms).Error != nil {
// 		 return map[uint]string{}
// 	 }
// 	 mMap := make(map[uint]string, 0)
// 	 for _, m := range ms {
// 	 	 mMap[m.ID] = m.Name
// 	 }
//	 return mMap
// }
