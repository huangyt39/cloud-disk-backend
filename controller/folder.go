package controller

import (
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/huangyt39/cloud-disk-backend/database"
    "github.com/huangyt39/cloud-disk-backend/models"
    "github.com/sirupsen/logrus"
    "net/http"
)

func GetFolders(c *gin.Context){
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
    foldersStr := string(foldersJson)
    fmt.Println(foldersStr)
    rawData := json.RawMessage(foldersStr)
    c.JSON(http.StatusOK, gin.H{
        "message" : "OK",
        "data" : rawData,
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

func GetFolder(c *gin.Context){
    //get files infomation
    var files []models.File
    db := database.DB.Where("folder = ?", c.Param("folder_name")).Find(&files)
    if err := db.Error; err != nil{
        logrus.Errorf("error on select folders, %s", err)
        c.JSON(http.StatusConflict, gin.H{
            "message" : "error",
        })
    }
    filesJson, err := json.Marshal(files)
    if err != nil{
        logrus.Errorf("error on 2json, %s", err)
    }
    filesStr := string(filesJson)
    rawData := json.RawMessage(filesStr)
    //get folder id
    var folder models.Folder
    db = database.DB.Where("name = ?", c.Param("folder_name")).Find(&folder)
    if err := db.Error; err != nil{
        logrus.Errorf("error on select folders, %s", err)
        c.JSON(http.StatusConflict, gin.H{
            "message" : "error",
        })
    }
    c.JSON(http.StatusOK, gin.H{
        "message" : "OK",
        "data" : gin.H{
            "files" : rawData,
            "id" : folder.ID,
            "name" : c.Param("folder_name"),
        },
    })
}

func UploadFolder(c *gin.Context){

}

func DeleteFolder(c *gin.Context){

}

