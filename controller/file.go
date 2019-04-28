package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/huangyt39/cloud-disk-backend/config"
    "github.com/huangyt39/cloud-disk-backend/database"
    "github.com/huangyt39/cloud-disk-backend/models"
    "github.com/huangyt39/cloud-disk-backend/utils"
    "github.com/sirupsen/logrus"
    "io"
    "net/http"
    "os"
    "path"
    "strconv"
)

func DownloadFile(c *gin.Context){
    cookie, _ := c.Request.Cookie("token")
    token := cookie.Value
    currentUser := database.GetCurrentUser(token)

    realfilename := utils.GenFilename(c.Param("folder_name"), c.Param("file_name"), currentUser)
    if !utils.PathExist(path.Join(config.SavePath, realfilename)){
        logrus.Errorf("error on find if file exit, file not exit")
        c.JSON(http.StatusNotFound, gin.H{
            "message" : "error",
        })
        return
    }
    // 设置浏览器是否为直接下载文件，且为浏览器指定下载文件的名字
    c.Header("Content-Disposition", `form-data; name="file"; filename=` + c.Param("file_name"))
    c.Header("Content-Type", "application/octet-stream")
    c.File(path.Join(config.SavePath, realfilename))
    c.JSON(http.StatusOK, gin.H{
        "message" : "ok",
    })
    return
}

func UploadFile(c *gin.Context){
    cookie, _ := c.Request.Cookie("token")
    token := cookie.Value
    currentUser := database.GetCurrentUser(token)
    //get folder id
    var folder models.Folder
    db := database.DB.Where("name = ? AND user_name = ?", c.Param("folder_name"), currentUser).Find(&folder)
    if err := db.Error; err != nil{
        logrus.Errorf("error on select folder, %s", err)
        c.JSON(http.StatusNotFound, gin.H{
            "message" : "error",
        })
        return
    }
    //get file infomation
    f, header, err := c.Request.FormFile("file")
    if err != nil{
        logrus.Errorf("error on get file, %s", err)
        c.JSON(http.StatusBadRequest, gin.H{
            "message" : "error",
        })
        return
    }
    filename := header.Filename
    //save file
    realfilename := utils.GenFilename(folder.Name, filename, currentUser)
    if utils.PathExist(path.Join(config.SavePath, realfilename)){
        logrus.Errorf("error on find if file exit")
        c.JSON(http.StatusConflict, gin.H{
            "message" : "error",
        })
        return
    }
    out, err := os.Create(path.Join(config.SavePath, realfilename))
    if err != nil{
        logrus.Errorf("error on save file, %s", err)
        c.JSON(http.StatusConflict, gin.H{
            "message" : "error",
        })
        return
    }
    _, err = io.Copy(out, f)
    if err != nil{
        logrus.Errorf("error on copy file, %s", err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "message" : "error",
        })
        return
    }
    defer out.Close()
    //create in db
    err = database.CreateFile(folder.ID, filename)
    if err != nil{
        logrus.Errorf("error on create file in db, %s", err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "message" : "error",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message" : "ok",
    })
    return
}

func DeleteFile(c *gin.Context){
    cookie, _ := c.Request.Cookie("token")
    token := cookie.Value
    currentUser := database.GetCurrentUser(token)

    realfilename := utils.GenFilename(c.Param("folder_name"), c.Param("file_name"), currentUser)
    if !utils.PathExist(path.Join(config.SavePath, realfilename)){
        logrus.Errorf("error on find if file exit, file not exit")
        c.JSON(http.StatusNotFound, gin.H{
            "message" : "error",
        })
        return
    }else{
        //get folder id
        folderId, err := database.GetFolderIDbyName(c.Param("folder_name"), currentUser)
        if err != nil{
            logrus.Errorf("error on get folderid by name %s", err)
            c.JSON(http.StatusNotFound, gin.H{
                "message" : "error",
            })
            return
        }
        if err := os.Remove(path.Join(config.SavePath, realfilename)); err != nil{
            logrus.Errorf("error on delete file in os, %s", err)
        }
        var file models.File
        db := database.DB.Where(&models.File{Filename: c.Param("file_name"), FolderId:folderId}).Delete(&file)
        if db.Error != nil{
            logrus.Error("error on delete file, %s", err)
            c.JSON(http.StatusInternalServerError, gin.H{
                "message" : "error",
            })
            return
        }else{
            c.JSON(http.StatusOK, gin.H{
                "message" : "ok",
            })
            return
        }
    }
}

func PatchSharetype(c *gin.Context){
    shareType := c.Query("shareType")
    //get folder id
    folderId := c.Param("folder_name")
    folderIdint, _ := strconv.Atoi(folderId)
    var file models.File
    db := database.DB.Where(&models.File{Filename: c.Param("file_name"), FolderId:folderIdint}).First(&file)
    if err := db.Error;err != nil{
        logrus.Errorf("error on find file, %s", err)
        c.JSON(http.StatusNotFound, gin.H{
            "message" : "error",
        })
        return
    }else {
        open_private_share, open_public_share := false, false
        if shareType == "private"{
            open_private_share = true
        }
        if shareType == "public"{
            open_public_share = true
        }
        db.Model(&file).Updates(map[string]interface{}{"OpenPublicShare":open_public_share, "OpenPrivateShare": open_private_share})
        if db.Error != nil{
            logrus.Errorf("error on update file sharetype, %s", err)
            c.JSON(http.StatusOK, gin.H{
                "message" : "ok",
            })
        }
    }
}