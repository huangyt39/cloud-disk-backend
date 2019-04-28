package database

import (
    "github.com/huangyt39/cloud-disk-backend/models"
    "github.com/sirupsen/logrus"
)

var LoginStatus map[string]string

func getLoginStatusFromDB(){
    LoginStatus = make(map[string]string)
    var users []models.User
    db := DB.Find(&users)
    if err := db.Error; err != nil{
        logrus.Errorf("error on get users %s", err)
    }else{
        for _, user := range users{
            LoginStatus[user.Token[:20]] = user.Name
        }
    }
}

func GetCurrentUser(token string) string {
    return LoginStatus[token[:20]]
}