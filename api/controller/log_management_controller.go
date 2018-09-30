package controller

import (
	"net/http"
	"peckergo/api/model"
	ginutils "peckergo/api/utils/gin"

	"github.com/gin-gonic/gin"
)

// AllLogManagementsGet get all logManagement
func AllLogManagementsGet(c *gin.Context) {
	// 分表注释下面两行代码
	meta := model.TableMetaFromQuery(c)
	ginutils.WriteGinJSON(c, http.StatusOK, model.AllLogManagements(meta))
	// 分表取消注释下面三行代码
	// meta := model.TableMetaFromQuery(c, "suffix")
	// suffix := c.Query("suffix")
	// ginutils.WriteGinJSON(c, http.StatusOK, model.AllLogManagements(meta, suffix))
}
