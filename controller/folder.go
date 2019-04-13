package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/huangyt39/cloud-disk-backend/database"
    "github.com/sirupsen/logrus"
    "net/http"
)

func GetFolder(c *gin.Context){

}

func CreateFolder(c *gin.Context){
    name := c.Query("name")
    if name != ""{
        err := database.CreateFolder(name)
        if err != nil{
            logrus.Errorf("error on create folder, %s", err)
            c.JSON(http.StatusConflict, gin.H{
                "message": "error",
            })
        }else{
            c.JSON(http.StatusCreated, gin.H{
                "message": "ok",
            })
        }
    }else{
        logrus.Error("error on create folder, name is nil")
        c.JSON(http.StatusConflict, gin.H{
            "message": "error",
        })
    }
}

func DownloadFolder(c *gin.Context){

}

func UploadFolder(c *gin.Context){

}

func DeleteFolder(c *gin.Context){

}

