package utils

import (
    "github.com/gin-gonic/gin"
    "github.com/huangyt39/cloud-disk-backend/config"
)

func verifyAuthToken(token string )bool{
    if token == config.Token{
        return true
    }
    return false
}

func VerifyAuthTokenMiddleWare() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
    }
}