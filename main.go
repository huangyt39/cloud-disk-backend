package main

import (
	"github.com/gin-gonic/gin"
	"github.com/huangyt39/cloud-disk-backend/database"
	"github.com/huangyt39/cloud-disk-backend/routers"
	"io"
	"os"
)

func main() {
	database.InitMysql()
	server()
	defer database.DB.Close()
}

func server(){
	r := gin.New()

	f , _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r.Use(gin.Logger())
	routers.LoadRouters(r)
	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}
