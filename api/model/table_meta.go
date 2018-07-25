package model

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// TableFilterMode 过滤模式
type TableFilterMode int

const (
	_ TableFilterMode = iota
	// Equal 等于
	Equal
	// GreaterThan 大于
	GreaterThan
	// GreaterThanOrEqual 大于等于
	GreaterThanOrEqual
	// LesserThan 小于
	LesserThan
	// LesserThanOrEqual 小于等于
	LesserThanOrEqual
	// Range 区间
	Range
)

// TableMeta 表单相关 meta
type TableMeta struct {
	Pagination TablePagination `json:"pagination"`
	Filter     []TableFilter   `json:"filter"`
	Order      []TableOrder    `json:"order"`
}

// TableFilter 表单过滤
type TableFilter struct {
	Field  string          `json:"field"`
	Mode   TableFilterMode `json:"model"`
	Params []interface{}   `json:"params"`
}

// TableOrder 表单排序
type TableOrder struct {
	Field string `json:"field"`
	Order string `json:"order"`
}

// TablePagination 表单分页
type TablePagination struct {
	Page    int  `json:"page"`
	PerPage int  `json:"size"`
	Total   uint `json:"total"`
}

// WrapMeta 包裹 meta 查询字段
func WrapMeta(meta TableMeta, db *gorm.DB) *gorm.DB {
	newDB := WrapFilter(meta.Filter, db)
	newDB = WrapPagination(meta.Pagination, newDB)
	newDB = WrapOrder(meta.Order, newDB)

	return newDB
}

// WrapPagination 包裹分页查询字段
func WrapPagination(p TablePagination, db *gorm.DB) *gorm.DB {
	newDB := db
	if p.PerPage != 0 {
		newDB = newDB.Limit(p.PerPage)
		if p.Page != -1 {
			newDB = newDB.Offset(p.PerPage * p.Page)
		}
	}
	return newDB
}

// WrapOrder 包裹排序查询字段
func WrapOrder(orders []TableOrder, db *gorm.DB) *gorm.DB {
	newDB := db
	for _, o := range orders {
		newDB = newDB.Order(fmt.Sprintf("%s %s", o.Field, o.Order))
	}
	return newDB
}

// WrapFilter 包裹过滤查询字段
func WrapFilter(filters []TableFilter, db *gorm.DB) *gorm.DB {
	newDB := db
	for _, f := range filters {
		switch f.Mode {
		case Equal:
			newDB = newDB.Where(fmt.Sprintf("%s = ?", f.Field), f.Params[0])
		case GreaterThan:
			newDB = newDB.Where(fmt.Sprintf("%s > ?", f.Field), f.Params[0])
		case GreaterThanOrEqual:
			newDB = newDB.Where(fmt.Sprintf("%s >= ?", f.Field), f.Params[0])
		case LesserThan:
			newDB = newDB.Where(fmt.Sprintf("%s < ?", f.Field), f.Params[0])
		case LesserThanOrEqual:
			newDB = newDB.Where(fmt.Sprintf("%s <= ?", f.Field), f.Params[0])
		case Range:
			if len(f.Params) >= 2 {
				newDB = newDB.Where(fmt.Sprintf("%s >= ?", f.Field), f.Params[0]).Where(fmt.Sprintf("%s < ?", f.Field), f.Params[1])
			}
		}
	}
	return newDB
}

// TableMetaFromQuery 把路由传参格式化成 TableMeta， 可添加不解析参数名 excludeQuerys
// http://restAPI?limit={limit}&page={page}&order=field[&order=-field]&field={[^^^|^^|___|__]value}&field={[value1, value2]}
func TableMetaFromQuery(c *gin.Context, excludeQuerys ...string) *TableMeta {
	queryKeys := QueryKeys(c)
	queryKeys = ReduceSlice(queryKeys, excludeQuerys)
	meta := &TableMeta{
		Pagination: TablePagination{},
		Filter:     make([]TableFilter, 0),
		Order:      make([]TableOrder, 0),
	}
	for _, key := range queryKeys {
		switch key {
		case "page":
			page, _ := strconv.ParseInt(c.Query(key), 10, 64)
			meta.Pagination.Page = int(page)
		case "limit":
			limit, _ := strconv.ParseInt(c.Query(key), 10, 64)
			meta.Pagination.PerPage = int(limit)
		case "order":
			meta.Order = ParseOrderQuerys(c.QueryArray(key))
		default:
			meta.Filter = AppendFilterQuery(meta.Filter, key, c.Query(key))
			fmt.Println(meta.Filter)
		}
	}
	return meta
}

// ParseOrderQuerys 把 order queries 格式化成 []TableOrder
func ParseOrderQuerys(orderQuerys []string) []TableOrder {
	orders := make([]TableOrder, 0)
	for _, q := range orderQuerys {
		order := TableOrder{}
		if strings.HasPrefix(q, "-") {
			order.Order = "desc"
			order.Field = q[1:]
		} else {
			order.Order = "asc"
			order.Field = q
		}
		orders = append(orders, order)
	}
	return orders
}

// AppendFilterQuery 解析添加 filter query
func AppendFilterQuery(filters []TableFilter, key string, val string) []TableFilter {
	if m, _ := regexp.MatchString("^\\[*.(,*.)+\\]$", val); m {
		vals := strings.Split(val[1:len(val)-1], ",")
		params := make([]interface{}, 0)
		for _, v := range vals {
			params = append(params, ParamTypeTransform(v))
		}
		filters = append(filters, TableFilter{
			Field:  key,
			Mode:   Range,
			Params: params,
		})
		fmt.Println(filters)
	} else {

		filter := TableFilter{
			Field: key,
		}
		if strings.HasPrefix(val, "^^^") {
			filter.Mode = GreaterThan
			filter.Params = []interface{}{ParamTypeTransform(val[3:])}
		} else if strings.HasPrefix(val, "^^") {
			filter.Mode = GreaterThanOrEqual
			filter.Params = []interface{}{ParamTypeTransform(val[2:])}
		} else if strings.HasPrefix(val, "___") {
			filter.Mode = LesserThan
			filter.Params = []interface{}{ParamTypeTransform(val[3:])}
		} else if strings.HasPrefix(val, "__") {
			filter.Mode = LesserThanOrEqual
			filter.Params = []interface{}{ParamTypeTransform(val[2:])}
		} else {
			filter.Mode = Equal
			filter.Params = []interface{}{ParamTypeTransform(val)}
		}
		filters = append(filters, filter)
	}
	return filters
}

// ParamTypeTransform 参数尝试转换 int 和 float
func ParamTypeTransform(param string) interface{} {
	if ret, err := strconv.ParseInt(param, 10, 64); err == nil {
		return ret
	} else if ret, err := strconv.ParseFloat(param, 64); err == nil {
		return ret
	}
	return param
}

// QueryKeys 获取所有 query 键值
func QueryKeys(c *gin.Context) []string {
	querys := c.Request.URL.Query()
	keys := make([]string, 0)
	for k := range querys {
		keys = append(keys, k)
	}
	return keys
}

// ReduceSlice source - target
func ReduceSlice(source []string, target []string) []string {
	for _, t := range target {
		for i, s := range source {
			if t == s {
				source = append(source[:i], source[i+1:]...)
				break
			}
		}
	}
	return source
}
