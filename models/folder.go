package models

type Folder struct{
    Name string `gorm:"type:varchar(64);primary_key:true"`
}