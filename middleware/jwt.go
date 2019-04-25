package middleware

import (
    "fmt"
    "github.com/gin-gonic/gin"
   "github.com/huangyt39/cloud-disk-backend/utils"
   "net/http"
   "time"
)

var (
    SUCCESS = 1
    INVALID_PARAMS = 2
    ERROR_AUTH_CHECK_TOKEN_FAIL = 3
    ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 4
)

func JWT() gin.HandlerFunc {
   return func(c *gin.Context) {
       fmt.Println("JWT middleware")
       var code int

       code = SUCCESS
       cookie, err := c.Request.Cookie("token")
       if err != nil{
           c.JSON(http.StatusUnauthorized, gin.H{
               "message" : "unauthorized",
           })
           c.Abort()
           return
       }
       token := cookie.Value
       if token == "" {
           code = INVALID_PARAMS
       } else {
           claims, err := utils.ParseToken(token)
           if err != nil {
               code = ERROR_AUTH_CHECK_TOKEN_FAIL
           } else if time.Now().Unix() > claims.ExpiresAt {
               code = ERROR_AUTH_CHECK_TOKEN_TIMEOUT
           }
       }

       if code != SUCCESS {
           c.JSON(http.StatusUnauthorized, gin.H{
               "code" : code,
               "message" : "unauthorized",
           })

           c.Abort()
           return
       }

       c.Next()
   }
}
