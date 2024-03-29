package controller

import (
	"{{projectName}}/api/model"
	ginutils "{{projectName}}/api/utils/gin"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// New{{ModelName}}Post create {{modelName}}
func New{{ModelName}}Post(c *gin.Context) {
	var {{modelName}} model.{{ModelName}}

	if err := ginutils.BindGinJSON(c, &{{modelName}}); err == nil {
		if err := model.New{{ModelName}}(&{{modelName}}); err == nil {
			ginutils.WriteGinJSON(c, http.StatusOK, gin.H{})
		} else {
			ginutils.WriteGinJSON(c, http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
		return
	}

	ginutils.WriteGinJSON(c, http.StatusBadRequest, gin.H{
		"msg": "param error!",
	})
}

// All{{ModelName}}sGet get all {{modelName}}
func All{{ModelName}}sGet(c *gin.Context) {
	// TODO 分表注释下面两行代码
	meta := model.TableMetaFromQuery(c)
	ginutils.WriteGinJSON(c, http.StatusOK, model.All{{ModelName}}s(meta))
	// TODO 分表取消注释下面三行代码
	// meta := model.TableMetaFromQuery(c, "suffix")
	// suffix := c.Query("suffix")
	// ginutils.WriteGinJSON(c, http.StatusOK, model.All{{ModelName}}s(meta, suffix))

}

// {{ModelName}}ByIDGet get {{modelName}} by id
func {{ModelName}}ByIDGet(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	log.Info("{{ModelName}}ByIDGet ", id)
	m := model.{{ModelName}}ByID(uint(id))
	ginutils.WriteGinJSON(c, http.StatusOK, m)
}

// Update{{ModelName}}Put 更新 {{ModelName}}
func Update{{ModelName}}Put(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	m := &model.{{ModelName}}{}

	if err := ginutils.BindGinJSON(c, m); err != nil {
		ginutils.WriteGinJSON(c, http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	m.ID = uint(id)

	if err := model.Save{{ModelName}}(m); err != nil {
		ginutils.WriteGinJSON(c, http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ginutils.WriteGinJSON(c, http.StatusOK, gin.H{})
}

// {{ModelName}}Delete 更新 {{ModelName}}
func {{ModelName}}Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if err := model.Delete{{ModelName}}(uint(id)); err != nil {
		ginutils.WriteGinJSON(c, http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ginutils.WriteGinJSON(c, http.StatusOK, gin.H{})
}

// TODO 为前端暴露 ID-Name 映射
// All{{ModelName}}IDNameMapGet 获得所有 {{ModelName}} ID-Name 映射
// func All{{ModelName}}IDNameMapGet(c *gin.Context) {
//	 mMap := model.All{{ModelName}}IDNameMap()
// 	 ginutils.WriteGinJSON(c, http.StatusOK, mMap)
// }

