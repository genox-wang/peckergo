package controller

import (
	"net/http"
	"peckergo/api/model"
	ginutils "peckergo/api/utils/gin"

	"github.com/gin-gonic/gin"
)

// AllDbtablessGet get all dbtables
func AllDbtablessGet(c *gin.Context) {
	meta := model.TableMetaFromQuery(c)
	ginutils.WriteGinJSON(c, http.StatusOK, model.AllDbtabless(meta))
}
