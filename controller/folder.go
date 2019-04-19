package controller

import (
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/huangyt39/cloud-disk-backend/database"
    "github.com/huangyt39/cloud-disk-backend/models"
    "github.com/sirupsen/logrus"
    "io/ioutil"
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
    var folder map[string]interface{}
    body, _ := ioutil.ReadAll(c.Request.Body)
    json.Unmarshal(body, &folder)
    if _, ok := folder["name"]; !ok{
        logrus.Error("error on get folder name")
        c.JSON(http.StatusBadRequest, gin.H{
        "message": "error",
        })
    }
    name := folder["name"].(string)
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
    //get folder id
    folderId, err := database.GetFolderIDbyName(c.Param("folder_name"))
    if err != nil{
        logrus.Errorf("error on get folderid by name %s", err)
        c.JSON(http.StatusNotFound, gin.H{
            "message" : "error",
        })
        return
    }
    //get files infomation
    var files []models.File
    db := database.DB.Where("folder_id = ?", folderId).Find(&files)
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
    c.JSON(http.StatusOK, gin.H{
        "message" : "OK",
        "data" : gin.H{
            "files" : rawData,
            "id" : folderId,
            "name" : c.Param("folder_name"),
        },
    })
}

func DeleteFolder(c *gin.Context){
    //get folder id
    folderId, err := database.GetFolderIDbyName(c.Param("folder_name"))
    if err != nil{
        logrus.Errorf("error on get folderid by name %s", err)
        c.JSON(http.StatusNotFound, gin.H{
            "message" : "error",
        })
        return
    }
    db := database.DB.Delete(&models.Folder{folderId, c.Param("folder_name")})
    if err := db.Error; err != nil{
        logrus.Errorf("error on delete folder, %s", err)
        c.JSON(http.StatusConflict, gin.H{
            "message" : "error",
        })
    }
    var files []models.File
    db = database.DB.Where("folder_id = ?", folderId).Find(&files)
    if err := db.Error; err != nil{
        logrus.Errorf("error on select files in folder, %s", err)
        c.JSON(http.StatusConflict, gin.H{
            "message" : "error",
        })
    }
    for _, file := range files {
        db = database.DB.Delete(&file)
        if err := db.Error; err != nil {
            logrus.Errorf("error on delete files in folder, %s", err)
            c.JSON(http.StatusConflict, gin.H{
                "message": "error",
            })
        }
    }
    c.JSON(http.StatusOK, gin.H{
        "message" : "OK",
    })
}

