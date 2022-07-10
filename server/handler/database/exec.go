package database

import (
	"github.com/crackeer/go-api-svc/container"
	ginhelper "github.com/crackeer/gopkg/gin"
	"github.com/crackeer/gopkg/storage/database"
	"github.com/crackeer/gopkg/util"
	"github.com/gin-gonic/gin"
)

// Exec ...
//  @param ctx
func Exec(ctx *gin.Context) {
	params := ginhelper.AllParams(ctx)
	sql := util.LoadMap(params).GetString("sql", "")
	if len(sql) < 1 {
		ginhelper.Failure(ctx, ginhelper.CodeDefaultError, "empty sql")
		return
	}
	err := database.ExecSQL(container.GetWriteDB(ctx.Param("database")), sql)
	if err != nil {
		ginhelper.Failure(ctx, ginhelper.CodeDefaultError, err.Error())
		return
	}
	ginhelper.Success(ctx, map[string]interface{}{
		"sql": sql,
	})
}
