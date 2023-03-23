package entities

import (
	"mime/multipart"
)

type BatchUploadObject struct {
	FileHeader *multipart.FileHeader
	Key        string
	FileName   string
}
