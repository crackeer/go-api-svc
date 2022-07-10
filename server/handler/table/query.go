package table

import (
	ginhelper "github.com/crackeer/gopkg/gin"
	"github.com/gin-gonic/gin"
)

// Query ...
//  @param ctx
func Query(ctx *gin.Context) {
	tableObj, err := getTableObject(ctx)
	if err != nil {
		ginhelper.Failure(ctx, -1, err.Error())
		return
	}
	params := ginhelper.AllParams(ctx)
	list := tableObj.Query(params, 100)
	ginhelper.Success(ctx, list)
}
