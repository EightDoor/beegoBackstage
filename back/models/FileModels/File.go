package FileModels

import (
	"beegoBackstage/commModel"
)

type File struct {
	commModel.BaseModel
	FileName    string `json:"fileName"`
	Url         string `json:"url"`
	FileSize    int64  `json:"fileSize"`
	StoragePath string `json:"storagePath"`
}

type FileBusiness struct {
	File *File `json:"file" orm:"null;rel(fk)"`
}
