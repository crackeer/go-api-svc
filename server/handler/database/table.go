package database

import (
	"github.com/crackeer/go-api-svc/container"
	ginhelper "github.com/crackeer/gopkg/gin"
	"github.com/crackeer/gopkg/storage/database"
	"github.com/gin-gonic/gin"
)

// Tables ...
//  @param ctx
func Tables(ctx *gin.Context) {
	list := database.GetSQLiteTables(container.GetReadDB(ctx.Param("database")))
	ginhelper.Success(ctx, list)
}
