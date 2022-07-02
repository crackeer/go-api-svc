package table

import (
	"github.com/crackeer/go-api-svc/container"
	"github.com/crackeer/gopkg/storage/table"
	"github.com/gin-gonic/gin"
)

func getTableObject(ctx *gin.Context) (*table.Table, error) {
	tableName := ctx.Param("table")
	database := ctx.Param("database")

	return table.NewTable(container.GetWriteDB(database), tableName)
}
