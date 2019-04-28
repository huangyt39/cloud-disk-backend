package models

type User struct{
    Name        string `gorm:"type:varchar(64);primary_key:true" json:"name"`
    Password    string `gorm:"type:varchar(64)"`
    Token       string `gorm:"type:varchar(256)"`
}