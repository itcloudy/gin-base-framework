package router

import (
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/gin-base-framework/common"
	"github.com/itcloudy/gin-base-framework/middles"
)

//setTemplate set template
func setTemplate(router *gin.Engine) {

	if common.FileUploadInfo.Storage == "local" {
		router.StaticFS(common.UPLOAD_FILE_URL, gin.Dir(common.FileUploadInfo.Path, false))
	}

	if common.ServerInfo.SystemStaticFilePath != "" {
		router.StaticFS(common.SYSTEM_STATIC_FILE_URL, gin.Dir(common.ServerInfo.SystemStaticFilePath, false))

	}

	return
}

//InitRouter router init
func InitRouter() *gin.Engine {

	router := gin.Default()

	router.Use(middles.RequestLanguages())
	router.Use(middles.Cors())
	router.Use(middles.VisitHistory())
	// set template
	setTemplate(router)
	// set middlewares
	//router.Use(middles.Visit())
	router.Use(middles.JwtAuthorize())
	router.Use(gin.Recovery())
	// add routers
	addRouter(router)
	return router
}
