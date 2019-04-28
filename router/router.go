package router

import (
	"github.com/gin-gonic/gin"
	"github.com/huangyt39/cloud-disk-backend/controller"
	"github.com/huangyt39/cloud-disk-backend/middleware"
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
	//router.GET("/s/:path", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", gin.H{
	//		"title": "Share page",
	//	})
	//})
	router.POST("/login", controller.Login)
	router.POST("/register", controller.Register)

	authorized := router.Group("/")
	authorized.Use(middleware.JWT())
	{
		router.GET("/auth", controller.Auth)
		router.GET("/folders", controller.GetFolders)
		router.POST("/folders", controller.CreateFolder)
		router.GET("/folders/:folder_name",  controller.GetFolder)
		router.POST("/folders/:folder_name", controller.UploadFile)
		router.DELETE("/folders/:folder_name",  controller.DeleteFolder)
		router.GET("/folders/:folder_name/:file_name", controller.DownloadFile)
		router.DELETE("/folders/:folder_name/:file_name", controller.DeleteFile)
		router.PATCH("/folders/:folder_name/:file_name", controller.PatchSharetype)
		router.GET("/share/:path", controller.SharePath)
	}
}
