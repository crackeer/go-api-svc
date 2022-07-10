package table

import (
	"fmt"
	"net/http"

	ginhelper "github.com/crackeer/gopkg/gin"
	"github.com/crackeer/gopkg/util"
	"github.com/gin-gonic/gin"
)

// Solo ...
//  @param ctx
func Solo(ctx *gin.Context) {
	tableObj, err := getTableObject(ctx)
	field := ctx.Param("field")
	if err != nil {
		ginhelper.Failure(ctx, -1, err.Error())
		return
	}
	params := ginhelper.AllParams(ctx)
	data, err := tableObj.Get(params)
	if err != nil {
		ginhelper.Failure(ctx, -1, err.Error())
		return
	}
	if len(data) < 1 {
		ginhelper.Failure(ctx, -1, "record not found")
		return
	}
	if _, exists := data[field]; !exists {
		ginhelper.Failure(ctx, -1, fmt.Sprintf("filed `%s` not exists", field))
		return
	}
	builder := util.LoadMap(data)
	contentType := builder.GetString("content_type", "text/plain")
	content := builder.GetString(field, "")
	ctx.Data(http.StatusOK, contentType, []byte(content))
}
