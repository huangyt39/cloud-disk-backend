package utils

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/huangyt39/cloud-disk-backend/config"
    "github.com/sirupsen/logrus"
    "net/http"
)

func VerifyAuthTokenMiddleWare() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("VerifyAuthToken")
        ok := false
        if token, err := c.Request.Cookie("token");err != nil{
            logrus.Errorf("error on get cookie, %s", err)
        }else{
            if verifyAuthToken(token.Value){
                ok = true
            }
        }
        if verifyAuthToken(c.Request.Header.Get("Authorization")){
            ok = true
        }
        if verifyAuthToken(c.Query("token")){
            ok = true
        }
        if !ok{
            c.JSON(http.StatusOK, gin.H{
                "message": "unauthorized",
            })
            c.Abort()
        }
        c.Next()
    }
}

func verifyAuthToken(token string )bool{
    if token == config.Token{
        return true
    }
    return false
}