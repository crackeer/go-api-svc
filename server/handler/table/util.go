package table

import (
	"github.com/crackeer/go-api-svc/container"
	"github.com/crackeer/gopkg/storage"
	"github.com/crackeer/gopkg/storage/table"
	"github.com/gin-gonic/gin"
)

func getTableObject(ctx *gin.Context) (*table.Table, error) {
	tableName := ctx.Param("table")
	database := ctx.Param("database")

	tableObject := &table.Table{
		DB:       container.GetWriteDB(database),
		Driver:   storage.DriverSQLite,
		Name:     tableName,
		PageSize: 20,
		OrderBy:  "id desc",
	}

	return tableObject, nil
}
