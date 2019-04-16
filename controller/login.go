package controller

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/huangyt39/cloud-disk-backend/config"
    "github.com/huangyt39/cloud-disk-backend/utils"
    "github.com/sirupsen/logrus"
    "io/ioutil"
    "net/http"
    "time"
)

type loginRequset struct {
    Email       string
    Password    string
}

func Login(c *gin.Context){
    body := c.Request.Body
    b, err := ioutil.ReadAll(body)
    if err != nil{
        logrus.Errorf("error on read body, %s", err)
    }
    var req loginRequset
    err = json.Unmarshal(b, &req)
    if err != nil{
        logrus.Errorf("error on json2struct, %s", err)
    }
    if req.Email != "" && req.Password != "" {
        if usernameMatchPassword(req.Email, req.Password) {
            http.SetCookie(c.Writer, &http.Cookie{
                Name:    "token",
                Value:   utils.GenToken(),
                Expires: time.Now().Add(30 * time.Minute),
                Path:    "/",
            })
            c.JSON(http.StatusOK, gin.H{
                "message": "ok",
                "token":   utils.GenToken(),
            })
            return
        }
    }
    c.JSON(http.StatusUnauthorized, gin.H{
        "message": "unauthorized",
    })
}

func usernameMatchPassword(username, password string) bool{
    if username == config.Username && password == config.Password{
        return true
    }
    return false
}