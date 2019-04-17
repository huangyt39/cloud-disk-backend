package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/huangyt39/cloud-disk-backend/controller"
	"github.com/huangyt39/cloud-disk-backend/utils"
	"net/http"
)

func LoadRouters(router *gin.Engine) {
	loadRouters(router)
}

func loadRouters(router *gin.Engine) {
	router.Use(utils.Cors())
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main page",
		})
	})
	router.GET("/s/:path", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Share page",
		})
	})
	router.POST("/login", controller.Login)
	router.GET("/auth", utils.VerifyAuthTokenMiddleWare(), controller.Auth)
	router.GET("/folders", utils.VerifyAuthTokenMiddleWare(), controller.GetFolders)
	router.POST("/folders", utils.VerifyAuthTokenMiddleWare(), controller.CreateFolder)
	router.GET("/folders/:folder_name", utils.VerifyAuthTokenMiddleWare(), controller.GetFolder)
	router.POST("/folders/:folder_name", utils.VerifyAuthTokenMiddleWare(), controller.UploadFolder)
	router.DELETE("/folders/:folder_name", utils.VerifyAuthTokenMiddleWare(), controller.DeleteFolder)
	router.GET("/folders/:folder_name/:file_name", utils.VerifyAuthTokenMiddleWare(), controller.DownloadFile)
	router.PATCH("/folders/:folder_name/:file_name", utils.VerifyAuthTokenMiddleWare(), controller.UploadFile)
	router.DELETE("/folders/:folder_name/:file_name", utils.VerifyAuthTokenMiddleWare(), controller.DeleteFile)
	router.GET("/share/:path", utils.VerifyAuthTokenMiddleWare(), controller.SharePath)
}
