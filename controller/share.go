package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/huangyt39/cloud-disk-backend/config"
    "github.com/huangyt39/cloud-disk-backend/database"
    "github.com/huangyt39/cloud-disk-backend/models"
    "github.com/huangyt39/cloud-disk-backend/utils"
    "github.com/sirupsen/logrus"
    "net/http"
    "path"
)

func SharePath(c *gin.Context){
    cookie, _ := c.Request.Cookie("token")
    token := cookie.Value

    isPublic, isPrivate := false, false
    var files []models.File
    db := database.DB.Where(&models.File{PublicShareUrl:c.Param("path")}).Find(&files)
    if err := db.Error;err != nil{
        logrus.Errorf("error on find file by url, %s", err)
    }
    if len(files) != 0{
        isPublic = true
    }else{
        db := database.DB.Where(&models.File{PrivateShareUrl:c.Param("path")}).Find(&files)
        if err := db.Error;err != nil{
            logrus.Errorf("error on find file by url, %s", err)
        }
        if len(files) != 0 {
            isPrivate = true
        }else{
            logrus.Errorf("can not find the file by path")
            c.JSON(http.StatusNotFound, gin.H{
                "message": "error",
            })
            return
        }
    }
    var folder models.Folder
    db = database.DB.Where("Id = ?", files[0].FolderId).First(&folder)
    if err := db.Error; err != nil{
        logrus.Errorf("error on find folder_name by folder_id %s", err)
    }
    realfilename := utils.GenFilename(folder.Name, files[0].Filename, folder.UserName)

    if !((isPublic && files[0].OpenPublicShare) || (isPrivate && files[0].OpenPrivateShare)){
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "error",
        })
        return
    }

    if c.Query("download") == "true"{
        if !utils.PathExist(path.Join(config.SavePath, realfilename)){
            logrus.Errorf("error on find if file exit, file not exit")
            c.JSON(http.StatusNotFound, gin.H{
                "message" : "error",
            })
            return
        }
        if isPublic || c.Query("password") == files[0].PrivateSharePassword{
            c.Header("Content-Disposition", `form-data; name="file"; filename=` + files[0].Filename)
            c.Header("Content-Type", "application/octet-stream")
            c.File(path.Join(config.SavePath, realfilename))
            c.JSON(http.StatusOK, gin.H{
                "message": "ok",
                "data" : gin.H{
                    "filename" : files[0].Filename,
                    "folder" : folder.Name,
                    "open_public_share": files[0].OpenPublicShare,
                    "open_private_share": files[0].OpenPrivateShare,
                    "token" : token,
                },
            })
            return
        }else{
            c.JSON(http.StatusNotFound, gin.H{
                "message" : "error",
            })
            return
        }
    }
}
