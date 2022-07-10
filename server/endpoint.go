package server

import (
	"fmt"

	"github.com/crackeer/go-api-svc/define"
	"github.com/crackeer/go-api-svc/server/handler/database"
	"github.com/crackeer/go-api-svc/server/handler/table"
	ginhelper "github.com/crackeer/gopkg/gin"
	"github.com/gin-gonic/gin"
)

func Run(config *define.AppConfig) error {

	router := gin.New()
	gin.SetMode(gin.DebugMode)
	dbRouter := router.Group("database/:database", ginhelper.DoResponseJSON())
	dbRouter.GET("query/:table", table.Query)
	dbRouter.GET("list/:table", table.List)
	dbRouter.POST("create/:table", table.Create)
	dbRouter.POST("update/:table", table.Update)
	dbRouter.POST("delete/:table", table.Delete)
	dbRouter.GET("distinct/:table", table.Distinct)
	dbRouter.POST("exec/sql", database.Exec)
	dbRouter.GET("tables", database.Tables)

	return router.Run(fmt.Sprintf(":%d", config.Port))
}
