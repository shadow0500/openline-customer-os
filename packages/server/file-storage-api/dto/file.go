package dto

type File struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Extension   string `json:"extension"`
	Mime        string `json:"mime"`
	MetadataUrl string `json:"previewUrl"`
	DownloadUrl string `json:"downloadUrl"`
}
