package model

import (
	"fmt"
	"peckergo/api/config"
)

// Dbtables Dbtables模型
type Dbtables struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Rows      uint   `json:"rows"`
	DataSize  uint   `json:"data_size"`
	IndexSize uint   `json:"index_size"`
	TotalSize uint   `json:"total_size"`
}

// TableDbtables 返回表单Dbtables数据模型
type TableDbtables struct {
	Data []*Dbtables `json:"data"`
	Meta *TableMeta  `json:"meta"`
}

// AllDbtabless 获取所有 Dbtabless
func AllDbtabless(meta *TableMeta) *TableDbtables {
	page := meta.Pagination.Page
	size := meta.Pagination.PerPage
	dbName := config.GetString("db.dbName")

	sql := fmt.Sprintf(`SELECT CONCAT(table_schema,'.',table_name) AS 'name',   
    table_rows AS 'rows',   
    data_length AS 'data_size',   
    index_length AS 'index_size' ,   
    data_length+index_length AS 'total_size'  
	FROM information_schema.TABLES   
	WHERE table_schema LIKE '%s' order by total_size desc limit %d offset %d`, dbName, size, page)

	dbtabless := make([]*Dbtables, 0)

	DB.Raw(sql).Scan(&dbtabless)

	sql = fmt.Sprintf(`SELECT COUNT(*) AS count,
		sum(table_rows) AS 'rows',   
		sum(data_length) AS 'data_size',   
		sum(index_length) AS 'index_size' ,   
		sum(data_length+index_length) AS 'total_size' 
	FROM information_schema.TABLES   
	WHERE table_schema LIKE '%s'`, dbName)

	type CountStruct struct {
		Count     uint `json:"count"`
		Rows      uint `json:"rows"`
		DataSize  uint `json:"data_size"`
		IndexSize uint `json:"index_size"`
		TotalSize uint `json:"total_size"`
	}

	countStruct := &CountStruct{}

	DB.Raw(sql).Scan(countStruct)

	meta.Pagination.Total = uint(countStruct.Count)

	dbtabless = append([]*Dbtables{
		&Dbtables{
			Name:      "Total",
			Rows:      countStruct.Rows,
			DataSize:  countStruct.DataSize,
			IndexSize: countStruct.IndexSize,
			TotalSize: countStruct.TotalSize,
		},
	}, dbtabless...)

	return &TableDbtables{
		Data: dbtabless,
		Meta: meta,
	}
}
