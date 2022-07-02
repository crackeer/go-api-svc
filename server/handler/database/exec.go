package database

import (
	"github.com/crackeer/go-api-svc/container"
	ginhelper "github.com/crackeer/gopkg/gin"
	"github.com/crackeer/gopkg/storage/database"
	"github.com/gin-gonic/gin"
)

// Exec ...
//  @param ctx
func Exec(ctx *gin.Context) {
	bytes, err := ctx.GetRawData()
	if err != nil {
		ginhelper.Failure(ctx, ginhelper.CodeDefaultError, err.Error())
		return
	}
	err = database.ExecSQL(container.GetWriteDB(ctx.Param("database")), string(bytes))
	if err != nil {
		ginhelper.Failure(ctx, ginhelper.CodeDefaultError, err.Error())
		return
	}
	ginhelper.Success(ctx, map[string]interface{}{
		"sql": string(bytes),
	})
}
