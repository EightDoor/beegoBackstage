package FileModels

import CommModels "beegoBackstage/commModels"

type File struct {
	CommModels.BaseModel
	FileName    string `json:"fileName"`
	Url         string `json:"url"`
	FileSize    int64  `json:"fileSize"`
	StoragePath string `json:"storagePath"`
}

type FileBusiness struct {
	File *File `json:"file" orm:"null;rel(fk)"`
}
