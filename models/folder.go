package models

type Folder struct{
    Name string `gorm:"type:varchar(64);unique_index"`
}
