package models

type Folder struct{
    ID   int `gorm:"AUTO_INCREMENT;primary_key:true"`
    Name string `gorm:"type:varchar(64)"`
}