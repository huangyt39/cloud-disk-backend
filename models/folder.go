package models

type Folder struct{
    ID          int `gorm:"AUTO_INCREMENT;primary_key:true" json:"id"`
    Name        string `gorm:"type:varchar(64)" json:"name"`
    UserName    string `json:"user_name"`
}