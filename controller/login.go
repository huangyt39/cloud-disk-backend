package controller

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/huangyt39/cloud-disk-backend/database"
    "github.com/huangyt39/cloud-disk-backend/models"
    "github.com/huangyt39/cloud-disk-backend/utils"
    "github.com/sirupsen/logrus"
    "io/ioutil"
    "net/http"
    "time"
)

type loginRequset struct {
	Email    string
	Password string
}

func Login(c *gin.Context) {
	body := c.Request.Body
	b, err := ioutil.ReadAll(body)
	if err != nil {
		logrus.Errorf("error on read body, %s", err)
	}
	var req loginRequset
	err = json.Unmarshal(b, &req)
	if err != nil {
		logrus.Errorf("error on json2struct, %s", err)
	}

	username, password := req.Email, req.Password
	if username == "" || password == ""{
        c.JSON(http.StatusUnauthorized, gin.H{
            "message": "unauthorized",
        })
        return
    }
	isExist := usernameMatchPassword(username, password)
	if isExist {
		token, err := utils.GenerateToken(username, password)
		if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "message": "unauthorized",
            })
            return
		} else {
			database.LoginStatus[token[:20]] = username
			database.DB.Model(models.User{}).Where("name = ?", username).Update("token", token)
            c.JSON(http.StatusOK, gin.H{
                "message": "ok",
                "token" : token,
            })
            http.SetCookie(c.Writer, &http.Cookie{
                Name:    "token",
                Value:   token,
                Expires: time.Now().Add(30 * time.Minute),
                Path:    "/",
            })
            return
		}

	} else {
        c.JSON(http.StatusUnauthorized, gin.H{
            "message": "unauthorized",
        })
        return
	}
}

func usernameMatchPassword(username, password string) bool {
	var users []models.User
	user := &models.User{
		username,
		password,
		"",
	}
	db := database.DB.Where(&user).Find(&users)
	if err := db.Error; err != nil {
		logrus.Errorf("error on find user, %s", err)
		return false
	} else {
		if len(users) == 0 {
			return false
		} else {
			return true
		}
	}
}
