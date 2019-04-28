package controller

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/huangyt39/cloud-disk-backend/database"
    "github.com/huangyt39/cloud-disk-backend/models"
    "github.com/sirupsen/logrus"
    "io/ioutil"
    "net/http"
)

type registerRequset struct {
    Email    string
    Password string
}

func Register(c *gin.Context){
    body := c.Request.Body
    b, err := ioutil.ReadAll(body)
    if err != nil {
        logrus.Errorf("error on read body, %s", err)
    }
    var req registerRequset
    err = json.Unmarshal(b, &req)
    if err != nil {
        logrus.Errorf("error on json2struct, %s", err)
    }

    username, password := req.Email, req.Password
    if username == "" || password == ""{
        logrus.Error("username or the password is null")
        c.JSON(http.StatusConflict, gin.H{
            "message": "error",
        })
        return
    }else{
        db := database.DB.Create(&models.User{username, password, ""})
        if err := db.Error; err != nil{
            logrus.Errorf("error on create the user", err)
            c.JSON(http.StatusConflict, gin.H{
                "message": "error",
            })
            return
        }else{
            c.JSON(http.StatusOK, gin.H{
                "message": "ok",
            })
            return
        }
    }
}
