package models

type File struct{
    Folder                  Folder
    Filename                string
    PublicShareUrl          string
    PrivateShareUrl         string
    PrivateSharePassword    string
    OpenPublicShare         bool
    OpenPrivateShare        bool
}