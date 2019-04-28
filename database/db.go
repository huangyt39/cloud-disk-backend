package database

import (
    "github.com/huangyt39/cloud-disk-backend/models"
    "github.com/huangyt39/cloud-disk-backend/utils"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/sirupsen/logrus"
)

var DB *gorm.DB

func InitMysql(){
    var err error
    DB, err = gorm.Open("mysql", "root:19971128hyt@/cloud_disk?charset=utf8&parseTime=True&loc=Local")
    if err != nil{
        logrus.Errorf("error on initmysql %s", err)
    }
    createIfTableNotExit()
    getLoginStatusFromDB()
}

func createIfTableNotExit(){
    if !DB.HasTable(&models.Folder{}){
        DB.CreateTable(&models.Folder{})
    }
    if err := DB.Error; err != nil{
        logrus.Errorf("error on create table folders, %s", err)
    }
    if !DB.HasTable(&models.File{}){
        DB.CreateTable(&models.File{})
    }
    if err := DB.Error; err != nil{
        logrus.Errorf("error on create table files, %s", err)
    }
    if !DB.HasTable(&models.User{}){
        DB.CreateTable(&models.User{})
    }
    if err := DB.Error; err != nil{
        logrus.Errorf("error on create table users, %s", err)
    }
}

func CreateFolder(name string, username string) error{
    newFolder := models.Folder{Name:name, UserName:username}
    err := DB.Create(&newFolder).Error
    if err != nil{
        return err
    }
    return nil
}

func CreateFile(folder_id int, filename string) error{
    newFile := models.File{
        FolderId:               folder_id,
        Filename:               filename,
        PublicShareUrl:         utils.GenerateUrl(1024),
        PrivateShareUrl:        utils.GenerateUrl(2048),
        PrivateSharePassword:   utils.GeneratePassword(),
        OpenPublicShare:        false,
        OpenPrivateShare:       false,
    }
    err := DB.Create(&newFile).Error
    if err != nil{
        return err
    }
    return nil
}

func GetFolderIDbyName(folder_name string, username string) (int , error){
    var folder models.Folder
    db := DB.Where("name = ? AND user_name = ?", folder_name, username).Find(&folder)
    if err := db.Error; err != nil{
        return 0, err
    }else{
        return folder.ID, nil
    }
}