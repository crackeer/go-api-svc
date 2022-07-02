package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadObject
//  @param ctx
func UploadObject(ctx *gin.Context) {
	header, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	tmpFile, err := header.Open()

	if err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	_, err = ioutil.ReadAll(tmpFile)
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

}

type share struct {
	ID       int64 `json:"id"`
	Duration int64 `json:"duration"`
}
