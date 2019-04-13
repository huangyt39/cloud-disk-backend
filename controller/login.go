package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/huangyt39/cloud-disk-backend/config"
    "github.com/huangyt39/cloud-disk-backend/utils"
    "net/http"
)

func Login(c *gin.Context){
    if c.Query("email") != "" && c.Query("email") != ""{
        if usernameMatchPassword(c.Query("email"), c.Query("password")){
            c.JSON(http.StatusOK, gin.H{
                "message": "ok",
                "token": utils.GenToken(),
            })
        }
    }else{
        c.JSON(http.StatusUnauthorized, gin.H{
            "message": "unauthorized",
        })
    }
}

func usernameMatchPassword(username, password string) bool{
    if username == config.Username && password == config.Password{
        return true
    }
    return false
}