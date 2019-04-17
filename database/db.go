package database

import (
    "github.com/huangyt39/cloud-disk-backend/models"
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
}

func CreateFolder(name string) error{
    newFolder := models.Folder{Name:name}
    err := DB.Create(&newFolder).Error
    if err != nil{
        return err
    }
    return nil
}