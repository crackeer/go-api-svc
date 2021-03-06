package table

import (
	"fmt"

	ginhelper "github.com/crackeer/gopkg/gin"
	"github.com/crackeer/gopkg/util"
	"github.com/gin-gonic/gin"
)

// Delete ...
//  @param ctx
func Delete(ctx *gin.Context) {
	tableObj, err := getTableObject(ctx)
	if err != nil {
		ginhelper.Failure(ctx, -1, err.Error())
		return
	}
	params := ginhelper.AllParams(ctx)

	primaryKey := tableObj.GetPrimaryKey()
	primaryKeyValue := util.LoadMap(params).GetString(primaryKey, "")
	if len(primaryKeyValue) < 1 {
		ginhelper.Failure(ctx, -1, fmt.Sprintf("primaryKey `%s` value not found", primaryKey))
		return
	}

	result := tableObj.Delete(primaryKey, primaryKeyValue)
	if result < 1 {
		ginhelper.Failure(ctx, -1, "not delete")
		return
	}
	ginhelper.Success(ctx, map[string]interface{}{
		"affect": result,
	})
}
