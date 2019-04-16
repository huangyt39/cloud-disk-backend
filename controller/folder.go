package controller

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/huangyt39/cloud-disk-backend/database"
    "github.com/huangyt39/cloud-disk-backend/models"
    "github.com/sirupsen/logrus"
    "net/http"
)

func GetFolder(c *gin.Context){
    var folders []models.Folder
    db := database.DB.Find(&folders)
    if err := db.Error; err != nil{
        logrus.Errorf("error on select folders, %s", err)
        c.JSON(http.StatusConflict, gin.H{
            "message" : "error",
        })
    }
    foldersJson, err := json.Marshal(folders)
    if err != nil{
        logrus.Errorf("error on 2json, %s", err)
    }
    c.JSON(http.StatusOK, gin.H{
        "message" : "ok",
        "data" : string(foldersJson),
    })
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

