package models

type File struct{
    ID                      int`gorm:"AUTO_INCREMENT;primary_key:true"`
    Folder                  Folder
    Filename                string
    PublicShareUrl          string
    PrivateShareUrl         string
    PrivateSharePassword    string
    OpenPublicShare         bool
    OpenPrivateShare        bool
}