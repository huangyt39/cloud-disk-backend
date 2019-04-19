package models

type File struct{
    ID                      int`gorm:"AUTO_INCREMENT;primary_key:true"`
    Folder                  Folder `gorm:"ForeignKey:FolderId"`
    FolderId                int
    Filename                string`json:"filename"`
    PublicShareUrl          string `json:"public_share_url"`
    PrivateShareUrl         string  `json:"private_share_url"`
    PrivateSharePassword    string  `json:"private_share_password"`
    OpenPublicShare         bool `json:"open_public_share"`
    OpenPrivateShare        bool `json:"open_private_share"`
}