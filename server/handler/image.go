package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RenderImage
//  @param ctx
func RenderImage(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.Data(http.StatusOK, "image/jpeg", []byte(id))
}
